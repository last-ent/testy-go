// Package commands is responsible for providing functions that will execute test commands on behalf of user.
// It also provides some functionality to make it easier for user to interact with the system.
package commands

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func executeCommand(command string) (string, error) {
	fmt.Println(command)
	out := &bytes.Buffer{}

	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out

	err := cmd.Run()
	return string(out.Bytes()), err
}

// ExecuteDefaultTests runs all non-skippable tests in the provided directory including all sub-directories.
func ExecuteDefaultTests(dir string) (string, error) {
	return executeCommand(fmt.Sprintf("go test -short -v %s/...", dir))
}

// ExecuteAllTests runs all tests (including skippable tests) in the provided directory including all sub-directories.
func ExecuteAllTests(dir string) (string, error) {
	return executeCommand(fmt.Sprintf("go test -v %s/...", dir))
}

// ExecuteAllTestsInNestedDir runs all tests in a selected directory including nested sub-directories.
func ExecuteAllTestsInNestedDir(dirMap map[int]string, dirIndex int) (string, error) {
	return executeCommand(fmt.Sprintf("go test -v %s/...", dirMap[dirIndex]))
}

// ExecuteDefaultTestsInNestedDir runs all non-skippable tests in a selected directory including nested sub-directories.
func ExecuteDefaultTestsInNestedDir(dirMap map[int]string, dirIndex int) (string, error) {
	return executeCommand(fmt.Sprintf("go test -short -v %s/...", dirMap[dirIndex]))
}

// ExecuteDefaultTestsInDir runs all tests only available in a selected directory.
func ExecuteDefaultTestsInDir(dirMap map[int]string, dirIndex int) (string, error) {
	return executeCommand("go test -short -v " + dirMap[dirIndex])
}

// ExecuteAllTestsInDir runs all tests only available in a selected directory.
func ExecuteAllTestsInDir(dirMap map[int]string, dirIndex int) (string, error) {
	return executeCommand("go test -v " + dirMap[dirIndex])
}

// ExecuteDefaultTestsInFile runs all non-skippable tests within a selected file.
// NOTE: This will only work for files that follow the conventions
// * filename.go
// * filename_test.go
func ExecuteDefaultTestsInFile(fileMap map[int]string, fileIndex int) (string, error) {
	testFile := fileMap[fileIndex]
	file := strings.Replace(testFile, "_test", "", 1)
	return executeCommand(fmt.Sprintf("go test -short -v %s %s", file, testFile))
}

// ExecuteAllTestsInFile runs all tests (including skippable tests) within a selected file.
// NOTE: This will only work for files that follow the conventions
// * filename.go
// * filename_test.go
func ExecuteAllTestsInFile(fileMap map[int]string, fileIndex int) (string, error) {
	testFile := fileMap[fileIndex]
	file := strings.Replace(testFile, "_test", "", 1)
	return executeCommand(fmt.Sprintf("go test -v %s %s", file, testFile))
}
