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
	Output   string
	ExitCode int
	Duration time.Duration
	Error    error
}

type Runner struct {
	dangerousPatterns []*regexp.Regexp
	workingDir        string
	sandboxDir        string
	isSandboxed       bool
}

func NewRunner() *Runner {
	// Create a temporary sandbox directory
	sandboxDir, err := os.MkdirTemp("", "tryoutshell-*")
	if err != nil {
		// Fallback to current directory if temp creation fails
		cwd, _ := os.Getwd()
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

	// Safety check
	if r.isDangerous(command) {
		return CommandResult{
			Output:   "❌ Command blocked for security reasons",
			Error:    fmt.Errorf("command blocked for safety"),
			Duration: time.Since(start),
			ExitCode: 1,
		}
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

	// Add COSIGN_PASSWORD for cosign commands to work non-interactively
	if strings.Contains(command, "cosign") {
		env = append(env, "COSIGN_PASSWORD=test123")
	}

	cmd.Env = env

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// Combine output
	output := stdout.String()
	if stderr.Len() > 0 {
		stderrStr := stderr.String()

		// Filter out the password prompt from stderr for cosign
		if strings.Contains(command, "cosign") {
			stderrStr = strings.ReplaceAll(stderrStr, "Enter password for private key: ", "")
			stderrStr = strings.TrimSpace(stderrStr)
		}

		// Check for common permission issues and provide helpful messages
		if strings.Contains(stderrStr, "Permission denied") ||
			strings.Contains(stderrStr, "Read-only file system") {
			output = r.enhancePermissionError(stderrStr, command)
		} else if stderrStr != "" {
			if output != "" {
				output += "\n"
			}
			output += stderrStr
		}
	}

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = 1
		}
	}

	return CommandResult{
		Output:   strings.TrimSpace(output),
		ExitCode: exitCode,
		Duration: time.Since(start),
		Error:    err,
	}
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

// Verify checks if command output matches expectations
func (r *Runner) Verify(result CommandResult, validation lessons_pkg.ValidationType) bool {
	switch validation.Type {
	case "regex":
		pattern := validation.Pattern
		if validation.CaseInsensitive {
			pattern = "(?i)" + pattern
		}
		re, err := regexp.Compile(pattern)
		if err != nil {
			return false
		}
		return re.MatchString(result.Output)

	case "substring":
		if validation.CaseInsensitive {
			return strings.Contains(
				strings.ToLower(result.Output),
				strings.ToLower(validation.Contains),
			)
		}
		return strings.Contains(result.Output, validation.Contains)

	case "exit_code":
		return result.ExitCode == validation.Expected

	case "file_exists":
		for _, file := range validation.Files {
			filePath := file
			if !filepath.IsAbs(file) {
				filePath = filepath.Join(r.workingDir, file)
			}

			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				return false
			}
		}
		return true

	case "output_contains":
		matchCount := 0
		for _, pattern := range validation.Patterns {
			if strings.Contains(result.Output, pattern) {
				matchCount++
				if validation.AnyMatch {
					return true
				}
			}
		}

		if validation.AllMatch {
			return matchCount == len(validation.Patterns)
		}

		return matchCount > 0

	default:
		return result.ExitCode == 0
	}
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
