package cmdx

import (
	"os"
	"path/filepath"
	"strconv"
)

// Config holds the configuration settings for the kcmd library.
type Config struct {
	PrintCommandEnabled bool   // Flag to enable or disable printing the command executed
	CommandDir          string // Directory to use to execute the commands
}

// DefaultConfig provides default settings for the library.
var DefaultConfig = Config{
	PrintCommandEnabled: getEnvAsBool("CMDX_PRINT_COMMAND_ENABLED", false), // Not print the command by default
	CommandDir:          getCommandDir(),
}

// SetConfig allows users to set custom configuration options.
func SetConfig(cfg Config) {
	DefaultConfig = cfg
}

// getCommandDir returns the directory for executing the Command.
// It first checks the CMDX_COMMAND_DIR environment variable.
// If not set, it defaults to the directory of the executable.
//
// Returns:
//
//	string: The base Command directory.
func getCommandDir() string {
	// Check if the KCMD_BASE_COMMAND_DIR environment variable is set
	commandDir := getEnv("CMDX_COMMAND_DIR", "")
	// If the environment variable is not set, use the directory of the executable
	// Also if the baseCommandDir is set to ".", use the directory of the executable
	if commandDir == "" {
		ex, err := os.Executable()
		if err != nil {
			panic(err) // Panic if the executable path cannot be determined
		}
		executablePath := filepath.Dir(ex)
		commandDir = executablePath
	}

	return commandDir
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the function returns its value.
// Otherwise, it returns the specified default value.
//
// Parameters:
//   - key: The name of the environment variable to look up.
//   - defaultVal: The value to return if the environment variable is not set.
//
// Returns:
//
//	string: The value of the environment variable or the default value if the variable is not set.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Same as getEnv but returns a boolean value.
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
