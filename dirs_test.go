package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMakeDirs(t *testing.T) {
	dt := &dirTree{
		"japan": dirTree{
			"nagoya": nil,
			"osaka":  nil,
			"tokyo": dirTree{
				"shibuya":  nil,
				"shinjuku": nil,
			},
		},
		"malaysia": dirTree{
			"kuala_lumpur": nil,
		},
		"singapore": nil,
	}
	wantDirs := []string{
		"japan",
		"japan/nagoya",
		"japan/osaka",
		"japan/tokyo",
		"japan/tokyo/shibuya",
		"japan/tokyo/shinjuku",
		"malaysia",
		"malaysia/kuala_lumpur",
		"singapore",
	}

	rootDir, err := ioutil.TempDir("", "TestMakeDirs")
	if err != nil {
		t.Fatalf("cannot create tempdir: %s", err)
	}
	defer os.RemoveAll(rootDir)

	if err := makeDirs(rootDir, dt); err != nil {
		t.Errorf("want no error, got %s", err)
	}

	for _, wd := range wantDirs {
		dir := filepath.Join(rootDir, wd)

		if s, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				t.Errorf("dir %q does not exist", dir)
			} else {
				t.Errorf("unexpected error at checking %q", dir)
			}
		} else {
			if !s.IsDir() {
				t.Errorf("%q is not a directory", dir)
			}
		}
	}
}
