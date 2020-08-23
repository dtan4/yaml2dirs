package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type dirTree map[string]interface{}

func parseYAMLFile(filename string) (*dirTree, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("cannot open the given YAML file: %w", err)
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	var data dirTree

	if err := d.Decode(&data); err != nil {
		return nil, fmt.Errorf("cannot parse the given YAML file: %w", err)
	}

	return &data, nil
}
