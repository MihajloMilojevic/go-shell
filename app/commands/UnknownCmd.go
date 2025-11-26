package commands

import (
	"fmt"
	"os"
)

type UnknownCommand struct {
	Cmd string
}

func NewUnknownCommand(cmd string, args []string) Command {
	return &UnknownCommand{Cmd: cmd}
}

func (c *UnknownCommand) Execute(IN, OUT *os.File) (bool, error) {
	_, err := fmt.Fprintf(OUT, "%s: command not found\n", c.Cmd)
	return false, err
}
