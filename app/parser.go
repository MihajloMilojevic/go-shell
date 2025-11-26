package main

import (
	"errors"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/commands/registries"

	_ "github.com/codecrafters-io/shell-starter-go/app/commands/builtins"
)

func Parse(cmd string) (commands.Command, error) {
	parts := splitCmd(cmd)
	if len(parts) == 0 {
		return nil, errors.New("empty command")
	}
	cmd = parts[0]
	args := parts[1:]
	if constructor, ok := registries.BuiltIns[cmd]; ok {
		return constructor(cmd, args), nil
	}
	return commands.NewUnknownCommand(cmd, args), nil
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
