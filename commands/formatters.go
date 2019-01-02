package commands

import (
	"fmt"
	"unicode/utf8"

	"githug.com/last-ent/testy-go/classifier"
)

func listDirs(meta *classifier.TestableDirMeta) string {
	return mapToString(listDirsAsMap(meta))
}

func mapToString(m map[int]string) string {
	strRepr := ""
	for index, item := range m {
		strRepr += fmt.Sprintf("\t%02d  ->  %s\n", index, item)
	}
	return strRepr
}

func prefixDir(dir string, path string) string {
	_, runeSize := utf8.DecodeRuneInString(path)
	return fmt.Sprintf("%s%s", dir, path[runeSize:])
}

func listFiles(meta *classifier.TestableDirMeta) string {
	return mapToString(listFilesAsMap(meta))
}

func listAll(meta *classifier.TestableDirMeta) string {
	dirsRepr := listDirs(meta)
	filesRepr := listFiles(meta)

	return fmt.Sprintf("DIRS:\n%s\n==========\n\nFILES:\n%s", dirsRepr, filesRepr)
}
