# Examples
This directory contains example usage of the `cmdx` library. 
Below are the details of each example provided.

## Basic
### Example 1 - `RunCommandPrintOutput`
- [Here](./run_ps_print_output.go) is an example to demonstrate how to use the `RunCommandPrintOutput` function to run a command and print the output.
- This example runs the `ps aux` command and prints the output to the console. 

### Example 2 - `RunCommandReturnOutput`
- [Here](./run_date_return_output.go) is an example to demonstrate how to use the `RunCommandWriteOutputToFile` function to run a command and return the output.
- This example runs the `date +%H:%M` command and save the result to the `out` variable.
- Then it prints the value of the `out` variable to the console.

### Example 3 - `RunCommandWriteOutputToFile`
- [Here](./run_ps_write_output_to_file.go) is an example to demonstrate how to use the `RunCommandWriteOutputToFile` function to run a command and write the output to a file.
- This example runs the `ps aux` command and writes the output to the `ps_output.txt` file.

### Example 4 - `RunCommandPrintOutputWithDirAndEnv` - `cmdDir`
- [Here](./run_ls_print_output_with_dir.go) is an example to demonstrate how to use the `RunCommandPrintOutputWithDirAndEnv` function to run a command from a specific directory.
- In this example we use set the `cmdDir` argument to run the `ls` command from the `/tmp` directory.

### Example 5 - `RunCommandPrintOutputWithDirAndEnv` - `cmdDir` and `envVars` 
- [Here](./run_kubeclt_get_nodes_with_env.go) is an example to demonstrate how to use the `RunCommandPrintOutputWithDirAndEnv` function to run a command from a specific directory and set additional environment variables.
- In this example we use the `envVars` argument to:  
- programmatically set the `KUBECONFIG` and `AWS_PROFILE` environment variables.
- Then run the `kubectl get nodes` command.
- More information about `cmdDir := cmdx.DefaultConfig.CommandDir` will be provided in the advanced section.
 
## Advanced 

### Example 1 - `RunCommandPrintOutput` - `CMDX_COMMAND_DIR` env variable
- [Here](./run_pwd_and_ls.go) is an example to demonstrate how to use the `CMDX_COMMAND_DIR` environment variable to control the directory where the command is executed.
- You can set the `CMDX_COMMAND_DIR` environment variable to control the directory where the command is executed.
- Note that the location of the go executable depends on how you run your code, via `go run` or `go build`. 
#### Running the example with `go run`
```shell
cd examples
export CMDX_COMMAND_DIR=$(pwd)
go run run_pwd_and_ls_print_output.go
or
export CMDX_COMMAND_DIR=/tmp
go run run_pwd_and_ls_print_output.go 
```
- When we use the `go run` a temporary go executable is created in an unpredictable location.
- So it makes sense to set the `CMDX_COMMAND_DIR` to a specific directory.
#### Running the example with `go build`
```shell
cd examples
go build run_pwd_and_ls_print_output.go
./run_pwd_and_ls_print_output 
```

### Example 2 - Custom Configuration
- [Here](./run_ls_custom_config.go) is an example to demonstrate how to use the `cmdx.Config` struct to configure the behavior of the library.
- In this example, we set a custom `cmdx.Config` struct to:
    1. Specify the directory where the command is executed.
    2. Enable printing the executed command.
```
config := cmdx.Config{
	PrintCommandEnabled: true,
	CommandDir:          "/tmp",
}
cmdx.SetConfig(config)
```