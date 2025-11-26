package registries

import (
	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

var BuiltIns = map[string]func(string, []string) commands.Command{}

func RegisterBuiltin(name string, factory func(string, []string) commands.Command) {
	BuiltIns[name] = factory
}
