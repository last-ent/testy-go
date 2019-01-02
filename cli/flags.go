package cli

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"
)

// ParseOpts contains all the parsed flags.
type ParseOpts struct {
	Dir string
}

// ParseFlags is used to parse all flags needed to run the program.
func ParseFlags(getDirFn func() (string, error)) (*ParseOpts, error) {
	var workDir string
	currDir, err := getCurrentDir(getDirFn)
	if err != nil {
		return nil, err
	}
	flag.StringVar(&workDir, "dir", currDir, "Directory to parse and find testable Go files.")
	flag.Parse()

	workDir, _ = filepath.Abs(workDir)
	fmt.Println("WELCOME TO TESTY GO! A TOOL FOR EASY TESTING.\nCWD: " + workDir)

	return &ParseOpts{workDir}, nil
}

func getCurrentDir(getDirFn func() (string, error)) (string, error) {
	dir, err := getDirFn()
	if err != nil {
		return "", fmt.Errorf("Error occurred: %s", err)
	}
	dir = strings.TrimSuffix(dir, "/")
	dir = strings.TrimSuffix(dir, "/...")

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("Error occurred: %s", err)
	}
	return absDir, nil
}
