package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {

	// Control the directory where the command will be executed
	// There are two ways to set the directory where the command will be executed
	// 1. way
	// Use the CMDX_COMMAND_DIR environment variable
	//This example is to demonstrate how to use the CMDX_COMMAND_DIR environment variable to set the directory where the command will be executed
	// Examples
	//export CMDX_COMMAND_DIR=/tmp
	// or
	//export CMDX_COMMAND_DIR=$(pwd)

	// 2. way
	//Use the Config struct as shown below
	/*
		config := cmdx.Config{
			CommandDir: "/tmp",
		}
		cmdx.SetConfig(config)
	*/
	print("Command directory: ")
	command := "pwd"
	err := cmdx.RunCommandPrintOutput(command)
	if err != nil {
		log.Fatalf("Error executing 'pwd' command: %v", err)
	}
	println("-")
	command = "ls"
	args := []string{"-la"}
	err = cmdx.RunCommandPrintOutput(command, args...)
	if err != nil {
		log.Fatalf("Error executing 'ls -la' command: %v", err)
	}
}
