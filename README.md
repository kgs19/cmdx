# cmdx
A Go library for executing and managing command-line operations with robust error handling and output management.

## Overview
The `cmdx` library simplifies command execution in Go applications. 
It provides functions to run commands, capture output, and manage errors effectively with structured error types `CommandError`. 
Additionally, `cmdx` includes a configurable `Config` struct, which allows users to customize execution settings, 
such as enabling command output printing and specifying the default directory for command execution.

## Installation
To install the library, run the following command:
```shell
go get -u github.com/kgs19/cmdx
```

## Usage

### Running Commands
Use `RunCommandPrintOutput` to execute a command and print the output:
```
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
```

### Handling Errors with `CommandError`
If a command fails, a `CommandError` instance is returned, providing detailed information:
```
package main

import (
	"fmt"
	"github.com/kgs19/cmdx"
	"log"
)

func main() {
	err := cmdx.RunCommandPrintOutput("command1.sh")
	if cmdErr, ok := err.(*cmdx.CommandError); ok {
		fmt.Printf("Command failed with exit code %d\n", cmdErr.ExitCode)
		fmt.Printf("Error message: %s\n", cmdErr.ErrorMsg)
	} else if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
}

```

### Usage Examples
View the [examples](examples/readme.md) directory for more examples of using the `cmdx` library.


## Configuration
The cmdx library includes configurable settings through the `Config` struct, allowing you to customize command execution:
  - `PrintCommandEnabled`: Set to true to enable printing each command before execution. This can be also controlled by the environment variable `CMDX_PRINT_COMMAND_ENABLED`.
  - `CommandDir `:Specify the directory for executing the commands. This can be also controlled by the environment variable `CMDX_COMMAND_DIR`.

Examples of updating configuration:  
    1. [example custom config](examples/readme.md#example-2---custom-configuration)  
    2. [example env variable](examples/readme.md#example-1---runcommandprintoutput---cmdx_command_dir-env-variable) 


## Contributing
Contributions are welcome! Feel free to submit issues or create pull requests.



