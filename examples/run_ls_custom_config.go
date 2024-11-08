package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	command := "ls"
	args := []string{"-la"}

	config := cmdx.Config{
		PrintCommandEnabled: true,
		CommandDir:          "/tmp",
	}
	cmdx.SetConfig(config)

	err := cmdx.RunCommandPrintOutput(command, args...)
	if err != nil {
		log.Fatalf("Error executing 'ls -la' command: %v", err)
	}
}
