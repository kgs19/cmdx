package main

import (
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	cmdDir := cmdx.DefaultConfig.CommandDir
	kubeconfigFilePath := "/tmp/kubeconfig"
	envVars := []string{
		"AWS_PROFILE=dev",
		"KUBECONFIG=" + kubeconfigFilePath,
	}

	command := "kubectl"
	args := []string{"get", "nodes"}

	err := cmdx.RunCommandPrintOutputWithDirAndEnv(command, cmdDir, envVars, args...)
	if err != nil {
		log.Fatalf("Error executing 'kubectl get nodes' command: %v", err)
	}
}
