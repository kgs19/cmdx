package cmdx

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// RunCommandPrintOutput executes a command and prints the output to os.Stdout.
//
// Parameters:
// - command: The command to be executed
// - args: Additional arguments to pass to the command
//
// Returns:
// - error: An error if the command fails, otherwise nil
func RunCommandPrintOutput(command string, args ...string) error {
	cmdDir := DefaultConfig.CommandDir
	return RunCommandPrintOutputWithDirAndEnv(command, cmdDir, nil, args...)
}

// RunCommandPrintOutputWithDirAndEnv runs a command with specified directory and environment variables.
// It prints the output to os.Stdout.
//
// Parameters:
// - command: The command to be executed
// - cmdDir: The directory in which to execute the command
// - envVars: A slice of additional environment variables to set for the command
// - args: Additional arguments to pass to the command
//
// Returns:
// - error: An error if the command fails, otherwise nil
func RunCommandPrintOutputWithDirAndEnv(command string, cmdDir string, envVars []string, args ...string) error {
	output := os.Stdout
	return runCommand(command, cmdDir, envVars, output, args...)
}

// RunCommandReturnOutput executes a command and returns the output as string.
//
// Parameters:
// - command: The command to be executed
// - args: Additional arguments to pass to the command
//
// Returns:
// - error: An error if the command fails, otherwise nil
func RunCommandReturnOutput(command string, args ...string) (string, error) {
	cmdDir := DefaultConfig.CommandDir
	return RunCommandReturnOutputWithDirAndEnv(command, cmdDir, nil, args...)
}

// RunCommandReturnOutputWithDirAndEnv runs a command with specified directory and environment variables.
// It returns the output as string.
//
// Parameters:
// - command: The command to be executed
// - cmdDir: The directory in which to execute the command
// - envVars: A slice of additional environment variables to set for the command
// - args: Additional arguments to pass to the command
//
// Returns:
// - string: The output of the command as a string
// - error: An error if the command fails, otherwise nil
func RunCommandReturnOutputWithDirAndEnv(command string, cmdDir string, envVars []string, args ...string) (string, error) {
	var output bytes.Buffer
	err := runCommand(command, cmdDir, envVars, &output, args...)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

// RunCommandWriteOutputToFile executes a command and writes the output to a specified file.
//
// Parameters:
// - command: The command to be executed
// - args: Additional arguments to pass to the command
//
// Returns:
// - error: An error if the command fails, otherwise nil
func RunCommandWriteOutputToFile(command string, filePath string, args ...string) error {
	cmdDir := DefaultConfig.CommandDir
	return RunCommandWriteOutputToFileWithDirAndEnv(command, cmdDir, nil, filePath, args...)
}

// RunCommandWriteOutputToFileWithDirAndEnv runs a command with specified directory and environment variables.
// It writes the output to a specified file.
//
// Parameters:
// - command: The command to be executed
// - cmdDir: The directory in which to execute the command
// - envVars: A slice of additional environment variables to set for the command
// - filePath: The path to the file where the command's output will be written
// - args: Additional arguments to pass to the command
//
// Returns:
// - error: An error if the command fails, otherwise nil
func RunCommandWriteOutputToFileWithDirAndEnv(command string, cmdDir string, envVars []string, filePath string, args ...string) error {
	// outFile, err := os.Create(fileAbsPath)
	//use append mode to append to the file if it exists
	outFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	//close the file when done.
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			fmt.Println("Error closing file", err)
		}
	}(outFile)

	//Set the output to be the file
	output := outFile
	return runCommand(command, cmdDir, envVars, output, args...)
}

// runCommand executes a command with specified directory and environment variables.
// It writes the output to the provided writer interface.
// It optionally logs the command being executed and the command directory, see DefaultConfig.PrintCommandEnabled.
//
// Parameters:
// - command: The command to be executed.
// - cmdDir: The directory in which to execute the command
// - envVars: A slice of additional environment variables to set for the command
// - output: An io.Writer where the command's standard output will be written.
// - args: Additional arguments to pass to the command.
//
// Returns:
// - error: An error if the command fails, otherwise nil.
//
// If the command fails,
// it captures the message printed in the standard error and exit code, and returns a custom error.
func runCommand(command string, cmdDir string, envVars []string, output io.Writer, args ...string) error {
	exitCode := 0
	var errb bytes.Buffer

	if DefaultConfig.PrintCommandEnabled {
		// Log the command details
		printCmd(command, cmdDir, output, args...)
	}

	// Set up the command with the provided directory and arguments
	cmd := exec.Command(command, args...)
	cmd.Dir = cmdDir

	// Set the environment variables from envVars
	setCmdEnvVars(cmd, envVars)

	// pipe the commands output to the applications
	// standard output
	cmd.Stdout = output
	cmd.Stderr = &errb
	err := cmd.Run()

	if err != nil {
		exitCode = 1
		stdErrorMsg := errb.String()

		// If no error message is captured in stderr, use the err.Error() instead
		if stdErrorMsg == "" {
			stdErrorMsg = err.Error()
		}

		var exitError *exec.ExitError
		if errors.As(err, &exitError) { // errors.As() -> function allows you to extract a specific error type from the error chain
			exitCode = exitError.ExitCode() //try to get actual cmd ExitCode
		}
		err := NewCommandError(stdErrorMsg, exitCode, cmdDir, command, args...)
		return err
	}
	return nil
}

func setCmdEnvVars(cmd *exec.Cmd, envVars []string) {
	cmd.Env = os.Environ()
	if envVars != nil && len(envVars) > 0 {
		for _, envVar := range envVars {
			cmd.Env = append(cmd.Env, envVar)
		}
	}
}

func commandWithArgs(command string, args ...string) string {
	return command + " " + strings.Join(args, " ")
}

func printCmd(command string, cmdDir string, output io.Writer, args ...string) {
	// For now do not print envVars may contain sensitive information
	cmd := commandWithArgs(command, args...)
	if cmdDir != "" {
		//Ignore error
		_, _ = fmt.Fprintf(output, "Execution directory: %s\n", cmdDir)
	}
	//print the command to output
	//Ignore error
	_, _ = fmt.Fprintf(output, "\nExecuting cmd: \n%s\n\n", cmd)
}
