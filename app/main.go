package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
		reader := bufio.NewReader(IN)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(OUT, "Error reading input:", err)
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		command = strings.TrimRight(line, "\r\n")
		cmd, err := Parse(command)
		if err != nil {
			fmt.Fprintln(OUT, "Error parsing command:", err)
			continue
		}
		shouldExit, err := cmd.Execute(IN, OUT)
		if err != nil {
			fmt.Fprintln(OUT, "Error executing command:", err)
		}
		if shouldExit {
			break
		}
	}
}
