package commands

import "os"

type Command interface {
	/// Return true if the shell should exit after executing this command
	Execute(IN, OUT *os.File) (bool, error)
}

type NilCommand struct{}

func (c *NilCommand) Execute(IN, OUT *os.File) (bool, error) {
	return false, nil
}

func NewNilCommand() *NilCommand {
	return &NilCommand{}
}
