package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	cmdDir := "c:\\Windows\\System32"
	command := "cmd"
	args := []string{"/C", "dir"} // bare format (no heading, summary, etc.)

	config := cmdx.Config{
		PrintCommandEnabled: true,
	}
	cmdx.SetConfig(config)

	out, err := cmdx.RunCommandReturnOutputWithDirAndEnv(command, cmdDir, nil, args...)
	if err != nil {
		log.Fatalf("Error executing 'dir' command: %v", err)
	}
	println(out)
}
