package cli

import (
	"fmt"
	"strings"
	"testing"
)

func getCurrentDirHelper(getDirFn func() (string, error), t *testing.T) {
	t.Helper()

	dir, err := getCurrentDir(getDirFn)
	if err != nil {
		t.Errorf("Unexpected error occurred: %s\n", err)
	}

	if !strings.HasSuffix(dir, "github.com/last-ent/testy-go/test_folder") {
		t.Errorf("Does not return expected format: %s\n", dir)
	}
}

func TestGetCurrentDir(t *testing.T) {
	for _, fn := range []func() (string, error){
		func() (string, error) {
			return "../test_folder", nil
		},
		func() (string, error) {
			return "../test_folder/", nil
		},
		func() (string, error) {
			return "../test_folder/...", nil
		},
	} {
		path, _ := fn()
		t.Run(path, func(t *testing.T) {
			getCurrentDirHelper(fn, t)
		})
	}
}

func TestParseFlags(t *testing.T) {
	for _, fn := range []func() (string, error){
		func() (string, error) {
			return "../test_folder/", nil
		},
		func() (string, error) {
			return "", fmt.Errorf("Expected")
		},
	} {
		opts, err := ParseFlags(fn)
		if err != nil {
			if err.Error() != "Error occurred: Expected" {
				t.Errorf("Unexepected error returned: %s", err)
			}
		} else if !strings.HasSuffix(opts.Dir, "github.com/last-ent/testy-go/test_folder") {
			t.Errorf("Does not return expected format: %s\n", opts.Dir)
		}
	}
}
