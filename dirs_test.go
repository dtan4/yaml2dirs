package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMakeDirs(t *testing.T) {
	testcases := map[string]struct {
		dt       *dirTree
		wantDirs []string
		wantErr  error
	}{
		"success": {
			dt: &dirTree{
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
			},
			wantDirs: []string{
				"japan",
				"japan/nagoya",
				"japan/osaka",
				"japan/tokyo",
				"japan/tokyo/shibuya",
				"japan/tokyo/shinjuku",
				"malaysia",
				"malaysia/kuala_lumpur",
				"singapore",
			},
			wantErr: nil,
		},
		"fail": {
			dt: &dirTree{
				"japan": "foo",
			},
			wantDirs: []string{},
			wantErr:  fmt.Errorf("is not a valid tree"),
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			rootDir, err := ioutil.TempDir("", "TestMakeDirs")
			if err != nil {
				t.Fatalf("cannot create tempdir: %s", err)
			}
			defer os.RemoveAll(rootDir)

			err = makeDirs(rootDir, tc.dt)

			if tc.wantErr == nil {
				for _, wd := range tc.wantDirs {
					if err != nil {
						t.Errorf("want no error, got %s", err)
					}

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
			} else {
				if err == nil {
					t.Errorf("want error %s, got no error", tc.wantErr)
				}

				if !strings.Contains(err.Error(), tc.wantErr.Error()) {
					t.Errorf("want error %q, got %q", tc.wantErr.Error(), err.Error())
				}
			}
		})
	}
}
