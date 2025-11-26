package builtins

import (
	"errors"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/commands/registries"
)

type TypeCommand struct {
	args []string
}

func NewTypeCommand(cmd string, args []string) commands.Command {
	return &TypeCommand{args: args}
}

func (c *TypeCommand) Execute(IN, OUT *os.File) (bool, error) {
	if len(c.args) == 0 {
		return false, errors.New("type: missing operand")
	}
	if _, ok := registries.BuiltIns[c.args[0]]; ok {
		OUT.WriteString(c.args[0] + " is a shell builtin\n")
	} else {
		OUT.WriteString(c.args[0] + ": not found\n")
	}
	return false, nil
}

func init() {
	registries.RegisterBuiltin("type", NewTypeCommand)
}
