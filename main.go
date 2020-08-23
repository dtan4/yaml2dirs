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
	filename := args[1]

	data, err := parseYAMLFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	fmt.Printf("%#v\n", data)

	return exitOK
}

func main() {
	os.Exit(realMain(os.Args))
}
