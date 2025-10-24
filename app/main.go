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
	fmt.Fprint(OUT, "$ ")
	var command string
	fmt.Fscanln(IN, &command)
	fmt.Fprintf(OUT, "%s: command not found\n", command)
}
