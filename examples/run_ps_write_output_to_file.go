package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	command := "ps"
	args := []string{"aux"}
	filePath := "ps_output.txt"

	err := cmdx.RunCommandWriteOutputToFile(command, filePath, args...)
	if err != nil {
		log.Fatalf("Error executing 'ps aux' command: %v", err)
	}
}
