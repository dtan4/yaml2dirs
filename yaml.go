package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func parseYAMLFile(filename string) (interface{}, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("cannot open the given YAML file: %w", err)
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	var data interface{}

	if err := d.Decode(&data); err != nil {
		return nil, fmt.Errorf("cannot parse the given YAML file: %w", err)
	}

	return data, nil
}
