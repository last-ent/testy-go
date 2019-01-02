package commands

import "github.com/last-ent/testy-go/classifier"

// ListDirsAsMap converts TestableDirMeta.Dirs into a map
func ListDirsAsMap(meta *classifier.TestableDirMeta) map[int]string {
	dirMap := map[int]string{}
	for index, dir := range meta.Dirs {
		dirMap[index+1] = prefixDir(meta.ParentDir, dir)
	}

	return dirMap
}

// ListFilesAsMap converts TestableDirMeta.Files into a map
func ListFilesAsMap(meta *classifier.TestableDirMeta) map[int]string {
	fileMap := map[int]string{}
	for index, file := range meta.Files {
		fileMap[index+1] = prefixDir(meta.ParentDir, file)
	}
	return fileMap
}
