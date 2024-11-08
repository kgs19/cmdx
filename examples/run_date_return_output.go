package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	command := "date"
	args := []string{"+%H:%M"}
	out, err := cmdx.RunCommandReturnOutput(command, args...)
	if err != nil {
		log.Fatalf("Error executing 'date' command: %v", err)
	}
	println("cmd output: " + out)
}
