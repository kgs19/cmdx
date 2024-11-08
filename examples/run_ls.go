package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	if err := cmdx.RunCommandPrintOutput("ls", "-l"); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}
