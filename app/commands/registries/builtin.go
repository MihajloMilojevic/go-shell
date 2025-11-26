package registries

import (
	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

var builtIns = map[string]func(string, []string) commands.Command{}

func RegisterBuiltin(name string, factory func(string, []string) commands.Command) {
	builtIns[name] = factory
}

func IsBuiltin(name string) bool {
	_, exists := builtIns[name]
	return exists
}

func GetBuiltInConstructor(name string) (func(string, []string) commands.Command, bool) {
	constructor, exists := builtIns[name]
	return constructor, exists
}
