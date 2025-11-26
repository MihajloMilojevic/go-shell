package builtins

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/commands/registries"
)

type ExitCommand struct{}

func NewExitCommand(cmd string, args []string) commands.Command {
	return &ExitCommand{}
}

func (c *ExitCommand) Execute(IN, OUT *os.File) (bool, error) {
	return true, nil
}

func init() {
	registries.RegisterBuiltin("exit", NewExitCommand)
}
