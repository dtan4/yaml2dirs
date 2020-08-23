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

	t, err := parseYAMLFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot parse the given YAML file: %w", err))
		return exitError
	}

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot get working directory: %w", err))
		return exitError
	}

	if err := makeDirs(rootDir, t); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot create directories: %w", err))
		return exitError
	}

	return exitOK
}

func main() {
	os.Exit(realMain(os.Args))
}
