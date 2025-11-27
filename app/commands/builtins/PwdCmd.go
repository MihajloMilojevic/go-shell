package builtins

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/commands/registries"
)

type PwdCommand struct{}

func NewPwdCommand(cmd string, args []string) commands.Command {
	return &PwdCommand{}
}

func (c *PwdCommand) Execute(IN, OUT *os.File) (bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return false, err
	}
	_, err = OUT.WriteString(dir + "\n")
	return false, err
}

func init() {
	registries.RegisterBuiltin("pwd", NewPwdCommand)
}
