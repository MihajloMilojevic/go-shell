package executables

import (
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

type ExecutableCommand struct {
	Name string
	Args []string
}

func NewExecutableCommand(name string, args []string) commands.Command {
	return &ExecutableCommand{
		Name: name,
		Args: args,
	}
}

func (ec *ExecutableCommand) Execute(IN, OUT *os.File) (bool, error) {
	execCmd := exec.Command(ec.Name, ec.Args...)
	execCmd.Stdin = IN
	execCmd.Stdout = OUT
	execCmd.Stderr = os.Stderr
	err := execCmd.Run()
	return false, err
}
