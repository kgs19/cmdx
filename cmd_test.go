package cmdx

import (
	"testing"
)

func TestRunCommandReturnOutputWithDirAndEnv(t *testing.T) {
	// Define test cases
	tests := []struct {
		command string
		cmdDir  string
		envVars []string
		args    []string
		wantErr bool
		wantOut string
	}{
		{"echo", "", nil, []string{"test"}, false, "test\n"},
		{"invalid_command", "", nil, nil, true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.command, func(t *testing.T) {
			output, err := RunCommandReturnOutputWithDirAndEnv(tt.command, tt.cmdDir, tt.envVars, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunCommandReturnOutput() error = %v, wantErr %v", err, tt.wantErr)
			}
			if output != tt.wantOut {
				t.Errorf("RunCommandReturnOutput() output = %v, wantOut %v", output, tt.wantOut)
			}
		})
	}
}
