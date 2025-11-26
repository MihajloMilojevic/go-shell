package commands

import (
	"fmt"
	"os"
)

type UnknownCommand struct {
	Cmd string
}

func (c *UnknownCommand) Execute(IN, OUT *os.File) (bool, error) {
	_, err := fmt.Fprintf(OUT, "%s: command not found\n", c.Cmd)
	return false, err
}

func NewUnknownCommand(cmd string) *UnknownCommand {
	return &UnknownCommand{Cmd: cmd}
}
