package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseYAMLFile(t *testing.T) {
	testcases := map[string]struct {
		filename string
		want     *dirTree
		wantErr  error
	}{
		"success": {
			filename: "testdata/dirs.yaml",
			want: &dirTree{
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
			wantErr: nil,
		},
		"file does not exist": {
			filename: "testdata/filedoesnotexist",
			want:     nil,
			wantErr:  fmt.Errorf("cannot open the given YAML file"),
		},
		"invalid YAML file": {
			filename: "testdata/invalid.yaml",
			want:     nil,
			wantErr:  fmt.Errorf("cannot parse the given YAML file"),
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := parseYAMLFile(tc.filename)

			if tc.wantErr == nil {
				if err != nil {
					t.Errorf("want no error, got %s", err)
				}

				if diff := cmp.Diff(*tc.want, *got); diff != "" {
					t.Errorf("-want, +got:\n%s", diff)
				}
			} else {
				if err == nil {
					t.Errorf("want error %s, got no error", tc.wantErr)
				}

				if !strings.Contains(err.Error(), tc.wantErr.Error()) {
					t.Errorf("want error %s, got %s", tc.wantErr.Error(), err.Error())
				}
			}
		})
	}
}
