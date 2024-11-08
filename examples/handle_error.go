package main

import (
	"fmt"
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	err := cmdx.RunCommandPrintOutput("somecommand", "--flag")
	if cmdErr, ok := err.(*cmdx.CommandError); ok {
		fmt.Printf("Command failed with exit code %d\n", cmdErr.ExitCode)
		fmt.Printf("Error message: %s\n", cmdErr.ErrorMsg)
	} else if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
}
