package commands

import (
	"reflect"
	"testing"

	"github.com/last-ent/testy-go/classifier"
)

func TestListDirsAsMap(t *testing.T) {
	meta := &classifier.TestableDirMeta{
		ParentDir: "/root",
		Files:     []string{},
		Dirs:      []string{"./dir1", "./dir2", "./dir3"},
	}

	actual := ListDirsAsMap(meta)
	expected := &OrderedMap{
		om: map[int]string{
			1: "/root/dir1",
			2: "/root/dir2",
			3: "/root/dir3",
		},
		order: []int{1, 2, 3},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%#v != %#v", actual, expected)
	}
}

func TestListFilesAsMap(t *testing.T) {
	meta := &classifier.TestableDirMeta{
		ParentDir: "/root",
		Files:     []string{"./dir1/file", "./dir2/file", "./dir3/file"},
		Dirs:      []string{},
	}

	actual := ListFilesAsMap(meta)
	expected := &OrderedMap{
		om: map[int]string{
			1: "/root/dir1/file",
			2: "/root/dir2/file",
			3: "/root/dir3/file",
		},
		order: []int{1, 2, 3},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%#v != %#v", actual, expected)
	}
}
