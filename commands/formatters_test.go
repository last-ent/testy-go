package commands

import (
	"strings"
	"testing"

	"github.com/last-ent/testy-go/classifier"
)

func TestMapToString(t *testing.T) {
	actual := mapToString(map[int]string{1: "1", 2: "2"})

	for _, expr := range []string{
		"\t01  ->  1\n",
		"\t02  ->  2\n",
	} {
		if !strings.Contains(actual, expr) {
			t.Errorf("Not found: %s in: \n%s", expr, actual)
		}
	}
}

func TestPrefixDir(t *testing.T) {
	actual := prefixDir("/root", "./path")
	expected := "/root/path"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}

func TestListDirs(t *testing.T) {
	meta := &classifier.TestableDirMeta{
		ParentDir: "/root",
		Files:     []string{},
		Dirs:      []string{"./dir1", "./dir2", "./dir3"},
	}

	actual := listDirs(meta)

	for _, expr := range []string{
		"\t01  ->  /root/dir1\n",
		"\t02  ->  /root/dir2\n",
		"\t03  ->  /root/dir3\n",
	} {
		if !strings.Contains(actual, expr) {
			t.Errorf("Not found: %s in: \n%s", expr, actual)
		}
	}
}

func TestListFiles(t *testing.T) {
	meta := &classifier.TestableDirMeta{
		ParentDir: "/root",
		Files:     []string{"./dir1/file", "./dir2/file", "./dir3/file"},
		Dirs:      []string{},
	}

	actual := listFiles(meta)

	for _, expr := range []string{
		"\t01  ->  /root/dir1/file\n",
		"\t02  ->  /root/dir2/file\n",
		"\t03  ->  /root/dir3/file\n",
	} {
		if !strings.Contains(actual, expr) {
			t.Errorf("Not found: %s in: \n%s", expr, actual)
		}
	}
}

func TestListAll(t *testing.T) {
	meta := &classifier.TestableDirMeta{
		ParentDir: "/root",
		Dirs:      []string{"./dir1", "./dir2", "./dir3"},
		Files:     []string{"./dir1/file", "./dir2/file", "./dir3/file"},
	}

	actual := listAll(meta)

	for _, expr := range []string{
		"\t01  ->  /root/dir1\n",
		"\t02  ->  /root/dir2\n",
		"\t03  ->  /root/dir3\n",
		"\t01  ->  /root/dir1/file\n",
		"\t02  ->  /root/dir2/file\n",
		"\t03  ->  /root/dir3/file\n",
	} {
		if !strings.Contains(actual, expr) {
			t.Errorf("Not found: %s in: \n%s", expr, actual)
		}
	}
}
