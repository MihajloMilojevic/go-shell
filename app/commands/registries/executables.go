package registries

import (
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/commands/executables"
)

func GetExecutablePath(name string) (string, bool) {
	return getExecutablePath(name)
}

func IsExecutable(name string) bool {
	_, found := getExecutablePath(name)
	return found
}

func GetExecutableConstructor(name string) (func(string, []string) commands.Command, bool) {
	if IsExecutable(name) {
		return executables.NewExecutableCommand, true
	}
	return nil, false
}

func getExecutablePath(name string) (string, bool) {
	for _, dir := range getPathDirs() {
		// Check if the directory exists and is a directory
		fileStat, err := os.Stat(dir)
		if err != nil || !fileStat.IsDir() {
			continue
		}
		// Check if the executable exists in this directory
		fullPath := dir + "/" + name
		fileStat, err = os.Stat(fullPath)
		if err != nil || fileStat.IsDir() {
			continue
		}
		// Check if the file is executable
		if fileStat.Mode()&0111 != 0 {
			return fullPath, true
		}
	}
	return "", false
}

func getPathDirs() []string {
	return strings.Split(os.Getenv("PATH"), ":")
}
