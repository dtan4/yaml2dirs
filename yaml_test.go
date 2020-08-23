package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseYAMLFile(t *testing.T) {
	filename := "testdata/dirs.yaml"
	want := map[string]interface{}{
		"japan": map[string]interface{}{
			"nagoya": nil,
			"osaka":  nil,
			"tokyo": map[string]interface{}{
				"shibuya":  nil,
				"shinjuku": nil,
			},
		},
		"malaysia": map[string]interface{}{
			"kuala_lumpur": nil,
		},
		"singapore": nil,
	}

	got, err := parseYAMLFile(filename)
	if err != nil {
		t.Errorf("want no error, got %s", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("-want, +got:\n%s", diff)
	}
}