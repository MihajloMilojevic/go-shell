package commands

import "os"

type EchoCommand struct {
	args []string
}

func NewEchoCommand(args []string) *EchoCommand {
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
