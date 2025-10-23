package runner

import (
	"bytes"
	"context"
	"fmt"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
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
}

func NewRunner() *Runner {
	return &Runner{
		dangerousPatterns: []*regexp.Regexp{
			regexp.MustCompile(`\bsudo\b`),
			regexp.MustCompile(`\brm\s+-rf\b`),
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
			Error:    fmt.Errorf("command blocked for safety reasons"),
			Duration: time.Since(start),
		}
	}

	// Set timeout
	if timeout == 0 {
		timeout = 30 // Default 30 seconds
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Execute command
	cmd := exec.CommandContext(ctx, "sh", "-c", command)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// Combine output
	output := stdout.String()
	if stderr.Len() > 0 {
		output += "\n" + stderr.String()
	}

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
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
		re := regexp.MustCompile(pattern)
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

	case "output_contains":
		for _, pattern := range validation.Patterns {
			matched := strings.Contains(result.Output, pattern)
			if validation.AnyMatch && matched {
				return true // Match any pattern
			}
			if validation.AllMatch && !matched {
				return false // Must match all patterns
			}
		}
		return validation.AllMatch // All matched if we got here

	default:
		return false
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
