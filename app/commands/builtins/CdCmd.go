package builtins

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/commands/registries"
)

type CdCommand struct {
	args []string
}

func NewCdCommand(cmd string, args []string) commands.Command {
	return &CdCommand{args: args}
}

func (c *CdCommand) Execute(IN, OUT *os.File) (bool, error) {
	if len(c.args) < 1 {
		return false, nil // No directory provided; do nothing
	}
	if c.args[0] == "~" {
		homeDir := os.Getenv("HOME")
		if homeDir != "" {
			c.args[0] = homeDir
		}
	}
	err := os.Chdir(c.args[0])
	if err != nil {
		OUT.WriteString("cd: " + c.args[0] + ": No such file or directory\n")
	}
	return false, nil
}

func init() {
	registries.RegisterBuiltin("cd", NewCdCommand)
}
