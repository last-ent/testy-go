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

// ExecuteDefaultTests runs all non-skippable tests in the provided package including all sub-packages.
func ExecuteDefaultTests(dir string) (string, error) {
	return executeCommand(fmt.Sprintf("go test -short -json -v %s/...", dir))
}

// ExecuteAllTests runs all tests (including skippable tests) in the provided package including all sub-packages.
func ExecuteAllTests(dir string) (string, error) {
	return executeCommand(fmt.Sprintf("go test -json -v %s/...", dir))
}

// ExecuteAllTestsInNestedDir runs all tests in a selected package including nested sub-packages.
func ExecuteAllTestsInNestedDir(dirMap *OrderedMap, dirIndex int) (string, error) {
	return executeCommand(fmt.Sprintf("go test -json -v %s/...", dirMap.Get(dirIndex)))
}

// ExecuteDefaultTestsInNestedDir runs all non-skippable tests in a selected package including nested sub-packages.
func ExecuteDefaultTestsInNestedDir(dirMap *OrderedMap, dirIndex int) (string, error) {
	return executeCommand(fmt.Sprintf("go test -short -json -v %s/...", dirMap.Get(dirIndex)))
}

// ExecuteDefaultTestsInDir runs all tests only available in a selected package.
func ExecuteDefaultTestsInDir(dirMap *OrderedMap, dirIndex int) (string, error) {
	return executeCommand("go test -short -json -v " + dirMap.Get(dirIndex))
}

// ExecuteAllTestsInDir runs all tests only available in a selected package.
func ExecuteAllTestsInDir(dirMap *OrderedMap, dirIndex int) (string, error) {
	return executeCommand("go test -json -v " + dirMap.Get(dirIndex))
}

// ExecuteDefaultTestsInFile runs all non-skippable tests within a selected file.
// NOTE: This will only work for files that follow the conventions
// * filename.go
// * filename_test.go
func ExecuteDefaultTestsInFile(fileMap *OrderedMap, fileIndex int) (string, error) {
	testFile := fileMap.Get(fileIndex)
	file := strings.Replace(testFile, "_test", "", 1)
	return executeCommand(fmt.Sprintf("go test -short -json -v %s %s", file, testFile))
}

// ExecuteAllTestsInFile runs all tests (including skippable tests) within a selected file.
// NOTE: This will only work for files that follow the conventions
// * filename.go
// * filename_test.go
func ExecuteAllTestsInFile(fileMap *OrderedMap, fileIndex int) (string, error) {
	testFile := fileMap.Get(fileIndex)
	file := strings.Replace(testFile, "_test", "", 1)
	return executeCommand(fmt.Sprintf("go test -json -v %s %s", file, testFile))
}
