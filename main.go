package main

import (
	"fmt"
	"os"
)

const (
	exitOK    = 0
	exitError = 1
)

func realMain(args []string) int {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "YAML file is required")
		return exitError
	}
	f := args[1]

	fmt.Println(f)

	return exitOK
}

func main() {
	os.Exit(realMain(os.Args))
}
