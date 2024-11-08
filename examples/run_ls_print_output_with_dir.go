package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	cmdDir := "/tmp"
	command := "ls"
	args := []string{"-la"}

	err := cmdx.RunCommandPrintOutputWithDirAndEnv(command, cmdDir, nil, args...)
	if err != nil {
		log.Fatalf("Error executing 'ls -la' command: %v", err)
	}
}
