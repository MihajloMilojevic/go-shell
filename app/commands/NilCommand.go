package commands

import (
	"os"
)

type NilCommand struct{}

func NewNilCommand(cmd string, args []string) Command {
	return &NilCommand{}
}

func (c *NilCommand) Execute(IN, OUT *os.File) (bool, error) {
	return false, nil
}
