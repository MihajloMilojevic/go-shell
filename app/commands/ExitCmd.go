package commands

import "os"

type ExitCommand struct{}

func (c *ExitCommand) Execute(IN, OUT *os.File) (bool, error) {
	return true, nil
}

func NewExitCommand() *ExitCommand {
	return &ExitCommand{}
}
