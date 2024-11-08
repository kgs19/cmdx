package cmdx

import "fmt"

// CommandError represents an error that occurred while executing a command.
// It includes the exit code, standard error message, and the command that was executed.
type CommandError struct {
	ExitCode int
	ErrorMsg string
	Command  string
	CmdDir   string
}

func (e *CommandError) Error() string {
	if e.CmdDir == "" {
		return fmt.Sprintf(
			"failed Command: \n%s\n"+
				"exit code: %d\n"+
				"error message: \n%s\n",
			e.Command, e.ExitCode, e.ErrorMsg)
	}
	return fmt.Sprintf(
		"failed Command: \n%s\n"+
			"exit code: %d\n"+
			"error message: \n%s\n"+
			"execution directory: %s\n",
		e.Command, e.ExitCode, e.ErrorMsg, e.CmdDir)
}

func NewCommandError(errorMsg string, exitCode int, cmdDir string, command string, args ...string) *CommandError {
	return &CommandError{
		ExitCode: exitCode,
		ErrorMsg: errorMsg,
		CmdDir:   cmdDir,
		Command:  commandWithArgs(command, args...),
	}
}
