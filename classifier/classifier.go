// Package classifier is responsible for parsing a given directory and provide all packages & files that can be tested.
package classifier

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TestableDirMeta provides information about all the files & dirs that can be tested
type TestableDirMeta struct {
	ParentDir string
	Dirs      []string
	Files     []string
}

func newTestableDirMeta(parentDir string, files []string, dirsMap map[string]bool) *TestableDirMeta {
	var dirs []string
	for dir := range dirsMap {
		dirs = append(dirs, dir)
	}
	return &TestableDirMeta{
		Dirs:      dirs,
		Files:     files,
		ParentDir: parentDir,
	}
}

// TraverseDir is responsible for traversing the folder structure
// at a given directory location to locate all possible packages that can be tested.
// The paths provided are relative to `dir`
func TraverseDir(dir string) (*TestableDirMeta, error) {
	var filesList []string
	dirsMap := map[string]bool{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err1 error) error {
		if err1 != nil {
			return err1
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "_test.go") {
			relPath, _ := filepath.Rel(dir, path)
			filesList = append(filesList, "./"+relPath)

			parentDir := filepath.Dir(path)
			relParentDir, _ := filepath.Rel(dir, parentDir)
			dirsMap["./"+relParentDir] = true
		}
		return nil
	})

	if err != nil {
		err1 := fmt.Errorf("error while traversing directory: %s.\n Error: %s", dir, err)
		return nil, err1
	}

	return newTestableDirMeta(dir, filesList, dirsMap), nil
}
