package main

import (
	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

func Parse(cmd string) (commands.Command, error) {
	parts := splitCmd(cmd)
	if len(parts) == 0 {
		return commands.NewNilCommand(), nil
	}
	cmd = parts[0]
	args := parts[1:]
	switch cmd {
	case "echo":
		return commands.NewEchoCommand(args), nil
	case "exit":
		return commands.NewExitCommand(), nil
	}
	return commands.NewUnknownCommand(cmd), nil
}

func splitCmd(cmd string) []string {
	var parts []string
	var cur []rune
	inQuotes := false

	for _, r := range cmd {
		switch r {
		case '"':
			inQuotes = !inQuotes
		case ' ', '\t':
			if inQuotes {
				cur = append(cur, r)
			} else {
				if len(cur) > 0 {
					parts = append(parts, string(cur))
					cur = cur[:0]
				}
				// skip extra separators
			}
		default:
			cur = append(cur, r)
		}
	}
	if len(cur) > 0 {
		parts = append(parts, string(cur))
	}
	return parts
}
