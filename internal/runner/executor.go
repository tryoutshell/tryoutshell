package runner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
)

type CommandResult struct {
	Command        string
	Output         string
	ExitCode       int
	Duration       time.Duration
	Error          error
	WorkingDir     string
	ValidationInfo ValidationResult
}
type ValidationResult struct {
	Passed         bool
	ValidationType string
	Expected       interface{}
	Actual         interface{}
	Details        []string
}

type Runner struct {
	dangerousPatterns []*regexp.Regexp
	workingDir        string
	sandboxDir        string
	isSandboxed       bool
}

func NewRunner() *Runner {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	// Create a temporary sandbox directory
	sandboxDir, err := os.MkdirTemp("", "tryoutshell-*")
	if err != nil {
		// Fallback to current directory if temp creation fails
		return &Runner{
			workingDir:        cwd,
			sandboxDir:        "",
			isSandboxed:       false,
			dangerousPatterns: compileDangerousPatterns(),
		}
	}

	// Make sandbox directory writable
	os.Chmod(sandboxDir, 0755)

	return &Runner{
		workingDir:        sandboxDir,
		sandboxDir:        sandboxDir,
		isSandboxed:       true,
		dangerousPatterns: compileDangerousPatterns(),
	}
}

func compileDangerousPatterns() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`\bsudo\b`),
		regexp.MustCompile(`\brm\s+-rf\s+/`),
		regexp.MustCompile(`\bchmod\s+777\b`),
		regexp.MustCompile(`>\s*/dev/sd[a-z]`),
		regexp.MustCompile(`curl.*\|.*sh`),
		regexp.MustCompile(`wget.*\|.*sh`),
		regexp.MustCompile(`:()\s*{\s*:\|:\s*&\s*};:`),
	}
}

// SetupLesson prepares the sandbox environment for a specific lesson
func (r *Runner) SetupLesson(lessonID string) error {
	if !r.isSandboxed {
		return fmt.Errorf("cannot setup lesson: not running in sandbox mode")
	}

	// Create common directories that lessons might need
	commonDirs := []string{
		"test",
		"temp",
		"data",
		"config",
		".config",
		"projects",
	}

	for _, dir := range commonDirs {
		dirPath := filepath.Join(r.workingDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create some sample files for practice
	sampleFiles := map[string]string{
		"README.md":      "# Welcome to TryOutShell\n\nThis is a practice environment.\n",
		"sample.txt":     "This is a sample text file.\n",
		"data/users.csv": "id,name,email\n1,Alice,alice@example.com\n2,Bob,bob@example.com\n",
	}

	for file, content := range sampleFiles {
		filePath := filepath.Join(r.workingDir, file)

		parentDir := filepath.Dir(filePath)
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			continue
		}

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			continue
		}
	}

	// Lesson-specific setup
	if strings.Contains(lessonID, "cosign") {
		return r.setupCosignLesson()
	}
	if strings.Contains(lessonID, "git") {
		return r.setupGitLesson()
	}

	return nil
}

// setupCosignLesson creates specific directories and files for cosign lesson
func (r *Runner) setupCosignLesson() error {
	// Create directories for keys and signed images
	dirs := []string{
		"keys",
		"images",
		".docker",
		".sigstore",
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(r.workingDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create cosign directory %s: %w", dir, err)
		}
	}

	return nil
}

// setupGitLesson initializes a git repository
func (r *Runner) setupGitLesson() error {
	repoPath := filepath.Join(r.workingDir, "test-repo")
	if err := os.MkdirAll(repoPath, 0755); err != nil {
		return err
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		return nil
	}

	readmePath := filepath.Join(repoPath, "README.md")
	os.WriteFile(readmePath, []byte("# Test Repository\n"), 0644)

	return nil
}

// Execute runs a shell command safely in the sandbox
func (r *Runner) Execute(command string, timeout int) CommandResult {
	start := time.Now()

	result := CommandResult{
		Command:    command,
		WorkingDir: r.workingDir,
	}
	// Safety check
	if r.isDangerous(command) {
		result.Output = "Command blocked for security reasons"
		result.Error = fmt.Errorf("command blocked for safety")
		result.Duration = time.Since(start)
		result.ExitCode = 126 // Command cannot execute
		return result
	}

	// Handle special debug commands
	if strings.HasPrefix(command, ":") {
		return r.handleDebugCommand(command)
	}

	// Set timeout
	if timeout == 0 {
		timeout = 30
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Execute command in working directory
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = r.workingDir

	// Set environment variables for the sandbox
	env := append(os.Environ(),
		"HOME="+r.workingDir,
		"TMPDIR="+filepath.Join(r.workingDir, "temp"),
	)

	// // Add COSIGN_PASSWORD for cosign commands to work non-interactively
	// if strings.Contains(command, "cosign") {
	// 	env = append(env, "COSIGN_PASSWORD=test123")
	// }

	cmd.Env = env

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// Combine output
	output := stdout.String()
	if stderr.Len() > 0 {
		if output != "" {
			output += "\n"
		}
		output += stderr.String()
	}

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else if ctx.Err() == context.DeadlineExceeded {
			exitCode = 124 // Timeout
			output += fmt.Sprintf("\n\nCommand timed out after %d seconds", timeout)
		} else {
			exitCode = 1
		}
	}

	result.Output = strings.TrimSpace(output)
	result.ExitCode = exitCode
	result.Duration = time.Since(start)
	result.Error = err

	return result
}

// handleDebugCommand handles special debug commands
func (r *Runner) handleDebugCommand(command string) CommandResult {
	result := CommandResult{
		Command:    command,
		WorkingDir: r.workingDir,
		Duration:   time.Millisecond,
		ExitCode:   0,
	}

	switch command {
	case ":pwd":
		result.Output = r.workingDir

	case ":ls":
		files, err := os.ReadDir(r.workingDir)
		if err != nil {
			result.Output = fmt.Sprintf("Error reading directory: %v", err)
			result.ExitCode = 1
			return result
		}

		var output strings.Builder
		output.WriteString(fmt.Sprintf("Directory: %s\n\n", r.workingDir))

		if len(files) == 0 {
			output.WriteString("(empty directory)\n")
		} else {
			output.WriteString("Files:\n")
			for _, file := range files {
				info, _ := file.Info()
				perms := info.Mode().String()
				size := info.Size()
				name := file.Name()

				if file.IsDir() {
					output.WriteString(fmt.Sprintf("  %s  %8d  %s/\n", perms, size, name))
				} else {
					output.WriteString(fmt.Sprintf("  %s  %8d  %s\n", perms, size, name))
				}
			}
		}
		result.Output = output.String()

	case ":state":
		result.Output = fmt.Sprintf("Working Directory: %s\n\n", r.workingDir)

		// Show current directory contents
		files, err := os.ReadDir(r.workingDir)
		if err != nil {
			result.Output += fmt.Sprintf("Error: %v\n", err)
		} else {
			result.Output += "Files in directory:\n"
			if len(files) == 0 {
				result.Output += "  (none)\n"
			} else {
				for _, file := range files {
					if file.IsDir() {
						result.Output += fmt.Sprintf("  📁 %s/\n", file.Name())
					} else {
						result.Output += fmt.Sprintf("  📄 %s\n", file.Name())
					}
				}
			}
		}

	case ":env":
		env := os.Environ()
		result.Output = strings.Join(env, "\n")

	default:
		result.Output = fmt.Sprintf("Unknown debug command: %s\n\nAvailable commands:\n  :pwd   - Show working directory\n  :ls    - List files\n  :env   - Show environment\n  :state - Show lesson state", command)
	}

	return result
}

// enhancePermissionError provides helpful hints for permission errors
func (r *Runner) enhancePermissionError(stderr, command string) string {
	hints := []string{
		"❌ Permission Error Detected",
		"",
		stderr,
		"",
		"💡 Hints:",
	}

	systemDirs := []string{"/etc", "/usr", "/var", "/sys", "/proc"}
	for _, dir := range systemDirs {
		if strings.Contains(command, dir) {
			hints = append(hints,
				"• You're trying to access a system directory ("+dir+")",
				"• Try using relative paths instead: ./your-file",
			)
			break
		}
	}

	if strings.Contains(command, "/home/") || strings.Contains(command, "/root/") {
		hints = append(hints,
			"• Avoid absolute paths like /home/user/",
			"• Use relative paths or current directory: ./file",
		)
	}

	if r.isSandboxed {
		testFile := filepath.Join(r.workingDir, ".write-test-"+fmt.Sprintf("%d", time.Now().Unix()))
		if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
			hints = append(hints,
				"• Sandbox directory is not writable",
				fmt.Sprintf("• Location: %s", r.workingDir),
				"• This is a bug - please report it!",
			)
		} else {
			os.Remove(testFile)
			hints = append(hints,
				"• Your sandbox IS writable",
				"• Try using relative paths: ./filename instead of /path/to/filename",
			)
		}
	}

	return strings.Join(hints, "\n")
}

// Verify checks if command output matches expectations WITH DETAILED FEEDBACK
func (r *Runner) Verify(result CommandResult, validation lessons_pkg.ValidationType) (CommandResult, bool) {
	validationResult := ValidationResult{
		ValidationType: validation.Type,
	}

	var passed bool

	switch validation.Type {
	case "regex":
		pattern := validation.Pattern
		if validation.CaseInsensitive {
			pattern = "(?i)" + pattern
		}
		re, err := regexp.Compile(pattern)
		if err != nil {
			validationResult.Details = append(validationResult.Details,
				fmt.Sprintf("Invalid regex pattern: %v", err))
			passed = false
		} else {
			passed = re.MatchString(result.Output)
			validationResult.Expected = pattern
			validationResult.Actual = result.Output

			if !passed {
				validationResult.Details = append(validationResult.Details,
					fmt.Sprintf("Output does not match regex pattern: %s", pattern))
			}
		}

	case "substring":
		expected := validation.Contains
		validationResult.Expected = expected

		if validation.CaseInsensitive {
			passed = strings.Contains(
				strings.ToLower(result.Output),
				strings.ToLower(expected),
			)
		} else {
			passed = strings.Contains(result.Output, expected)
		}

		if !passed {
			validationResult.Details = append(validationResult.Details,
				fmt.Sprintf("Expected substring '%s' not found in output", expected))
			validationResult.Actual = result.Output
		}

	case "exit_code":
		expected := validation.Expected
		actual := result.ExitCode
		passed = actual == expected

		validationResult.Expected = expected
		validationResult.Actual = actual

		if !passed {
			validationResult.Details = append(validationResult.Details,
				fmt.Sprintf("Expected exit code %d, got %d", expected, actual))
		}

	case "file_exists":
		validationResult.Expected = validation.Files
		var foundFiles []string
		var missingFiles []string

		for _, file := range validation.Files {
			fullPath := filepath.Join(r.workingDir, file)
			if _, err := os.Stat(fullPath); err == nil {
				foundFiles = append(foundFiles, file)
			} else {
				missingFiles = append(missingFiles, file)
			}
		}

		passed = len(missingFiles) == 0
		validationResult.Actual = foundFiles

		if !passed {
			validationResult.Details = append(validationResult.Details,
				fmt.Sprintf("Missing files: %v", missingFiles))
			validationResult.Details = append(validationResult.Details,
				fmt.Sprintf("Found files: %v", foundFiles))
		}

	case "output_contains":
		validationResult.Expected = validation.Patterns
		var matchedPatterns []string
		var unmatchedPatterns []string

		for _, pattern := range validation.Patterns {
			if strings.Contains(result.Output, pattern) {
				matchedPatterns = append(matchedPatterns, pattern)
				if validation.AnyMatch {
					passed = true
					break
				}
			} else {
				unmatchedPatterns = append(unmatchedPatterns, pattern)
			}
		}

		if validation.AllMatch {
			passed = len(unmatchedPatterns) == 0
		}

		validationResult.Actual = matchedPatterns

		if !passed {
			if validation.AnyMatch {
				validationResult.Details = append(validationResult.Details,
					"None of the expected patterns found in output")
			} else {
				validationResult.Details = append(validationResult.Details,
					fmt.Sprintf("Missing patterns: %v", unmatchedPatterns))
			}
		}

	default:
		// No validation specified, check exit code
		passed = result.ExitCode == 0
		validationResult.ValidationType = "exit_code"
		validationResult.Expected = 0
		validationResult.Actual = result.ExitCode

		if !passed {
			validationResult.Details = append(validationResult.Details,
				fmt.Sprintf("Command failed with exit code %d", result.ExitCode))
		}
	}

	validationResult.Passed = passed
	result.ValidationInfo = validationResult

	return result, passed
}

// isDangerous checks if command contains dangerous patterns
func (r *Runner) isDangerous(command string) bool {
	for _, pattern := range r.dangerousPatterns {
		if pattern.MatchString(command) {
			return true
		}
	}
	return false
}

// GetWorkingDir returns the current working directory
func (r *Runner) GetWorkingDir() string {
	return r.workingDir
}

// IsSandboxed returns whether the runner is operating in a sandbox
func (r *Runner) IsSandboxed() bool {
	return r.isSandboxed
}

// Cleanup removes the sandbox directory and all its contents
func (r *Runner) Cleanup() error {
	if !r.isSandboxed || r.sandboxDir == "" {
		return nil
	}

	if !strings.Contains(r.sandboxDir, "tryoutshell-") {
		return fmt.Errorf("refusing to remove directory that doesn't look like a sandbox: %s", r.sandboxDir)
	}

	return os.RemoveAll(r.sandboxDir)
}

// ListFiles returns a list of files in the sandbox (for debugging)
func (r *Runner) ListFiles() ([]string, error) {
	var files []string

	err := filepath.Walk(r.workingDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(r.workingDir, path)
		if err != nil {
			relPath = path
		}

		if !info.IsDir() {
			files = append(files, relPath)
		}
		return nil
	})

	return files, err
}

func (r *Runner) SetWorkingDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dir)
	}
	r.workingDir = dir
	return nil
}
