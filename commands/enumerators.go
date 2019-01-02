package commands

import "github.com/last-ent/testy-go/classifier"

func listDirsAsMap(meta *classifier.TestableDirMeta) map[int]string {
	dirMap := map[int]string{}
	for index, dir := range meta.Dirs {
		dirMap[index+1] = prefixDir(meta.ParentDir, dir)
	}

	return dirMap
}

func listFilesAsMap(meta *classifier.TestableDirMeta) map[int]string {
	fileMap := map[int]string{}
	for index, file := range meta.Files {
		fileMap[index+1] = prefixDir(meta.ParentDir, file)
	}
	return fileMap
}
