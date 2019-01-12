package commands

import (
	"fmt"
	"strings"

	"github.com/last-ent/testy-go/classifier"
)

// ListDirs returns string representation of all testable packages under `dir`
func ListDirs(meta *classifier.TestableDirMeta) string {
	return mapToString(ListDirsAsMap(meta))
}

func mapToString(om *OrderedMap) string {
	strRepr := ""
	for _, item := range om.Iterator() {
		strRepr += fmt.Sprintf("\t%02d  ->  %s\n", item.key, item.value)
	}
	return strRepr
}

func prefixDir(dir string, path string) string {
	path = strings.TrimPrefix(path, ".")
	return fmt.Sprintf("%s%s", dir, path)
}

// ListFiles returns string representation of all testable files under `dir`
func ListFiles(meta *classifier.TestableDirMeta) string {
	return mapToString(ListFilesAsMap(meta))
}

// ListAll returns string representation of all testable files & packages under `dir`
func ListAll(meta *classifier.TestableDirMeta) string {
	dirsRepr := ListDirs(meta)
	filesRepr := ListFiles(meta)

	return fmt.Sprintf("DIRS:\n%s\n==========\n\nFILES:\n%s", dirsRepr, filesRepr)
}
