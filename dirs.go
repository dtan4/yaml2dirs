package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func makeDirs(rootDir string, t *dirTree) error {
	for d, c := range *t {
		dir := filepath.Join(rootDir, d)

		if err := os.Mkdir(dir, 0755); err != nil {
			return fmt.Errorf("cannot create directory %q, %w", dir, err)
		}

		if c == nil {
			continue
		}

		ct, ok := c.(dirTree)
		if !ok {
			return fmt.Errorf("children of %q is not a valid tree", dir)
		}

		if err := makeDirs(dir, &ct); err != nil {
			return fmt.Errorf("cannot create directories under %q", dir)
		}
	}

	return nil
}
