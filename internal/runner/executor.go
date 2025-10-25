package runner

import (
	"bytes"
	"context"
	"fmt"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
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
}

func NewRunner() *Runner {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}

	return &Runner{
		workingDir: cwd,
		dangerousPatterns: []*regexp.Regexp{
			regexp.MustCompile(`\bsudo\b`),
			regexp.MustCompile(`\brm\s+-rf\s+/`), // Only block rm -rf /
			regexp.MustCompile(`\bchmod\s+777\b`),
			regexp.MustCompile(`>\s*/dev/sd[a-z]`),
			regexp.MustCompile(`curl.*\|.*sh`),
			regexp.MustCompile(`wget.*\|.*sh`),
			regexp.MustCompile(`:()\s*{\s*:\|:\s*&\s*};:`), // Fork bomb
		},
	}
}

// Execute runs a shell command safely
func (r *Runner) Execute(command string, timeout int) CommandResult {
	start := time.Now()

	// Safety check
	if r.isDangerous(command) {
		return CommandResult{
			Output:   "Command blocked for security reasons",
			Error:    fmt.Errorf("command blocked for safety"),
			Duration: time.Since(start),
			ExitCode: 1,
		}
	}

	// Set timeout
	if timeout == 0 {
		timeout = 30 // Default 30 seconds
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Execute command in working directory
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = r.workingDir

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
			if _, err := os.Stat(file); os.IsNotExist(err) {
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
					return true // Match any pattern
				}
			}
		}

		if validation.AllMatch {
			return matchCount == len(validation.Patterns)
		}

		return matchCount > 0

	default:
		// If no validation type specified, consider success if exit code is 0
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
