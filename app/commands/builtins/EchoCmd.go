package builtins

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/commands/registries"
)

type EchoCommand struct {
	args []string
}

func NewEchoCommand(cmd string, args []string) commands.Command {
	return &EchoCommand{args: args}
}

func (c *EchoCommand) Execute(IN, OUT *os.File) (bool, error) {
	for i, arg := range c.args {
		if i > 0 {
			OUT.WriteString(" ")
		}
		OUT.WriteString(arg)
	}
	OUT.WriteString("\n")
	return false, nil
}

func init() {
	registries.RegisterBuiltin("echo", NewEchoCommand)
}
