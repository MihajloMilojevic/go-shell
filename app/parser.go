package main

import "github.com/codecrafters-io/shell-starter-go/app/commands"

func Parse(cmd string) commands.Command {
	if cmd == "exit" {
		return commands.NewExitCommand()
	}
	return commands.NewUnknownCommand(cmd)
}
