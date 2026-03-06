package main

import (
	"fmt"
	"github.com/tryoutshell/tryoutshell/internal/runner"
)

func main() {
	fmt.Println("=== Sandbox Debug Test ===")
	fmt.Println()

	// Create runner
	r := runner.NewRunner()
	defer r.Cleanup()

	// Check if sandboxed
	fmt.Printf("Is Sandboxed: %v\n", r.IsSandboxed())
	fmt.Printf("Working Dir: %s\n\n", r.GetWorkingDir())

	// Try to create a file
	fmt.Println("Testing file creation...")
	testFile := "test-write.txt"

	result := r.Execute(fmt.Sprintf("echo 'Hello World' > %s", testFile), 5)
	fmt.Printf("Command: echo 'Hello World' > %s\n", testFile)
	fmt.Printf("Exit Code: %d\n", result.ExitCode)
	fmt.Printf("Output: %s\n", result.Output)
	if result.Error != nil {
		fmt.Printf("Error: %v\n", result.Error)
	}
	fmt.Println()

	// Try to read it back
	fmt.Println("Testing file read...")
	result2 := r.Execute("cat test-write.txt", 5)
	fmt.Printf("Command: cat test-write.txt\n")
	fmt.Printf("Exit Code: %d\n", result2.ExitCode)
	fmt.Printf("Output: %s\n", result2.Output)
	if result2.Error != nil {
		fmt.Printf("Error: %v\n", result2.Error)
	}
	fmt.Println()

	// List files
	fmt.Println("Listing files in sandbox...")
	result3 := r.Execute("ls -la", 5)
	fmt.Printf("Output:\n%s\n", result3.Output)

	// Try the actual cosign command
	fmt.Println("\n=== Testing Cosign Command ===")
	result4 := r.Execute("cosign generate-key-pair", 5)
	fmt.Printf("Command: cosign generate-key-pair\n")
	fmt.Printf("Exit Code: %d\n", result4.ExitCode)
	fmt.Printf("Output: %s\n", result4.Output)
	if result4.Error != nil {
		fmt.Printf("Error: %v\n", result4.Error)
	}

	fmt.Println("\n=== Test Complete ===")
}
