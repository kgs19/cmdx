package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	command := "ps"
	args := []string{"aux"}
	err := cmdx.RunCommandPrintOutput(command, args...)
	if err != nil {
		log.Fatalf("Error executing 'ps aux' command: %v", err)
	}
}
