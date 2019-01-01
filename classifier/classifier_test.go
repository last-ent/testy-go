package classifier

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTestDirMeta(t *testing.T) {
	testCases := []struct {
		Name     string
		Files    []string
		Dirs     map[string]bool
		Expected TestableDirMeta
	}{
		{
			Name:  "both files",
			Files: []string{"a", "b", "c"},
			Dirs: map[string]bool{
				"d1": true,
				"d2": false,
			},
			Expected: TestableDirMeta{ParentDir: "", Files: []string{"a", "b", "c"}, Dirs: []string{"d1", "d2"}},
		},
		{
			Name:     "only files",
			Files:    []string{"a", "b", "c"},
			Dirs:     map[string]bool{},
			Expected: TestableDirMeta{ParentDir: "", Files: []string{"a", "b", "c"}, Dirs: []string{}},
		},
		{
			Name:  "only dirs",
			Files: []string{},
			Dirs: map[string]bool{
				"d1": true,
				"d2": false,
			},
			Expected: TestableDirMeta{ParentDir: "", Files: []string{}, Dirs: []string{"d1", "d2"}},
		},
		{
			Name:     "empty dir",
			Files:    []string{},
			Dirs:     map[string]bool{},
			Expected: TestableDirMeta{ParentDir: "", Files: []string{}, Dirs: []string{}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := newTestableDirMeta("", tc.Files, tc.Dirs)
			expected := tc.Expected
			if !(reflect.DeepEqual(actual.Dirs, expected.Dirs) || reflect.DeepEqual(actual.Files, expected.Files)) {
				t.Error(fmt.Sprintf("%s != %s", actual, tc.Expected))
			}
		})
	}
}

func TestTraverseDir(t *testing.T) {
	dir := "../test_folder"

	expected := TestableDirMeta{
		ParentDir: dir,
		Files: []string{
			"./level1/file1_test.go",
			"./level1/file2_test.go",
			"./level2/file3_test.go",
			"./level2/sub-level/file4_test.go",
			"./level2/sub-level/file5_test.go",
			"./level3/sub-level/file7_test.go",
			"./level3/sub-level/sub-level2/file8_test.go",
		},
		Dirs: []string{
			"./level1",
			"./level2",
			"./level2/sub-level",
			"./level3/sub-level/sub-level2",
		},
	}

	actual, err := TraverseDir(dir)
	if err != nil {
		t.Error(err)
	}
	if !(reflect.DeepEqual(actual.Dirs, expected.Dirs) || reflect.DeepEqual(actual.Files, expected.Files) || actual.ParentDir == dir) {
		t.Error(fmt.Sprintf("%s != %s", actual, expected))
	}
}
