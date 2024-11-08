package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

// example usage of the kcmd package
func main() {
	cmdDir := "c:\\Windows\\System32"
	command := "cmd"
	args := []string{"/C", "dir"} // bare format (no heading, summary, etc.)

	config := cmdx.Config{
		PrintCommandEnabled: true,
	}
	cmdx.SetConfig(config)

	err := cmdx.RunCommandPrintOutputWithDirAndEnv(command, cmdDir, nil, args...)
	if err != nil {
		log.Fatalf("Error executing 'dir' command: %v", err)
	}
}
