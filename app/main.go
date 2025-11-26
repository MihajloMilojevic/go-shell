package main

import (
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	var IN = os.Stdin
	var OUT = os.Stdout
	for {
		fmt.Fprint(OUT, "$ ")
		var command string
		_, err := fmt.Fscanln(IN, &command)
		if err != nil {
			fmt.Fprintln(OUT, "Error reading input:", err)
			continue
		}
		cmd := Parse(command)
		shouldExit, err := cmd.Execute(IN, OUT)
		if err != nil {
			fmt.Fprintln(OUT, "Error executing command:", err)
		}
		if shouldExit {
			break
		}
	}
}
