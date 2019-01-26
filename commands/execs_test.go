package commands

import (
	"fmt"
	"strings"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	actual, err := executeCommand("echo 'Hello, World!'")
	if err != nil {
		t.Errorf("Error while executing tests: %s\n%s\n", err, actual)
		return
	}

	if actual != "Hello, World!\n" {
		t.Errorf("%s != Hello, World!\n", actual)
	}
}

func TestExecuteDefaultTests(t *testing.T) {
	actual, err := ExecuteDefaultTests("../test_folder")
	if err != nil {
		t.Errorf("Error while executing tests: %s\n%s\n", err, actual)
		return
	}

	expected := []struct {
		Name string
		Expr string
	}{
		{"TestFunc1 should pass", "--- PASS: TestFunc1"},
		{"TestFile2 should be skipped", "--- SKIP: TestFile2"},
		{"TestFunc3 should pass", "--- PASS: TestFunc3"},
		{"TestFunc4 should pass", "--- PASS: TestFunc4"},
		{"TestFile5 should be skipped", "--- SKIP: TestFile5"},
		{"TestFile7 should be skipped", "--- SKIP: TestFile7"},
		{"TestFunc8 should pass", "--- PASS: TestFunc8"},
	}

	for _, expr := range expected {
		t.Run(expr.Name, func(t *testing.T) {
			if !strings.Contains(actual, expr.Expr) {
				t.Errorf("Unexpected result:\n%s\n", actual)
			}
		})
	}
}

func TestExecuteAllTests(t *testing.T) {
	actual, _ := ExecuteAllTests("../test_folder")

	expected := []struct {
		Name string
		Expr string
	}{
		{"TestFunc1 should pass", "--- PASS: TestFunc1"},
		{"TestFile2 should fail", "--- FAIL: TestFile2"},
		{"TestFunc3 should pass", "--- PASS: TestFunc3"},
		{"TestFunc4 should pass", "--- PASS: TestFunc4"},
		{"TestFile5 should fail", "--- FAIL: TestFile5"},
		{"TestFile7 should fail", "--- FAIL: TestFile7"},
		{"TestFunc8 should pass", "--- PASS: TestFunc8"},
	}

	for _, expr := range expected {
		t.Run(expr.Name, func(t *testing.T) {
			if !strings.Contains(actual, expr.Expr) {
				t.Errorf("Unexpected result:\n%s\n", actual)
			}
		})
	}
}

func TestExecuteAllTestsInNestedDir(t *testing.T) {
	dirs := &OrderedMap{
		om: map[int]string{
			1: "../test_folder/level1",
			2: "../test_folder/level2",
		},
		order: []int{1, 2},
	}

	testCases := []struct {
		Index   int
		Name    string
		Results []string
		ErrMsg  string
	}{
		{
			1,
			"Simple Nesting",
			[]string{
				"-- PASS: TestFunc1",
				"-- FAIL: TestFile2",
			},
			"Simple nesting failed! Not found:",
		},
		{
			2,
			"Deep Nesting",
			[]string{
				"-- PASS: TestFunc3",
				"-- PASS: TestFunc4",
				"-- FAIL: TestFile5",
			},
			"Deep nesting failed! Not found:",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual, _ := ExecuteAllTestsInNestedDir(dirs, tc.Index)
			for _, expr := range tc.Results {
				if !strings.Contains(actual, expr) {
					t.Error(tc.ErrMsg + expr)
				}
			}
		})
	}
}

func TestExecuteDefaultTestsInNestedDir(t *testing.T) {
	dirs := &OrderedMap{
		om: map[int]string{
			1: "../test_folder/level1",
			2: "../test_folder/level2",
		},
		order: []int{1, 2},
	}

	testCases := []struct {
		Index   int
		Name    string
		Results []string
		ErrMsg  string
	}{
		{
			1,
			"Simple Nesting",
			[]string{
				"-- PASS: TestFunc1",
				"-- SKIP: TestFile2",
			},
			"Simple nesting failed! Not found:",
		},
		{
			2,
			"Deep Nesting",
			[]string{
				"-- PASS: TestFunc3",
				"-- PASS: TestFunc4",
				"-- SKIP: TestFile5",
			},
			"Deep nesting failed! Not found:",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual, _ := ExecuteDefaultTestsInNestedDir(dirs, tc.Index)
			for _, expr := range tc.Results {
				if !strings.Contains(actual, expr) {
					t.Error(tc.ErrMsg + expr)
				}
			}
		})
	}
}

func TestExecuteDefaultTestsInDir(t *testing.T) {
	dirs := &OrderedMap{
		om: map[int]string{
			1: "../test_folder/level2",
			2: "../test_folder/level3/sub-level",
			3: "../test_folder/level3/sub-level/sub-level2",
		},
		order: []int{1, 2, 3},
	}

	testCases := []struct {
		Index   int
		Name    string
		ErrMsg  string
		Results map[string]bool
	}{
		{
			Index: 1,
			Name:  "Level 2 Nesting",
			Results: map[string]bool{
				"TestFunc4":          false,
				"TestFile5":          false,
				"-- PASS: TestFunc3": true,
			},
			ErrMsg: "Level 2 nesting failed!",
		},
		{
			Index: 2,
			Name:  "Level 3.1 Nesting (Don't skip)",
			Results: map[string]bool{
				"-- SKIP: TestFile7": true,
			},
			ErrMsg: "Level 3.1 nesting failed!",
		},
		{
			Index: 3,
			Name:  "Level 3.2 Nesting",
			Results: map[string]bool{
				"TestFile7":          false,
				"-- PASS: TestFunc8": true,
			},
			ErrMsg: "Level 3.2 nesting failed!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual, _ := ExecuteDefaultTestsInDir(dirs, tc.Index)
			for expr, result := range tc.Results {
				if strings.Contains(actual, expr) != result {
					t.Errorf("%s Result contains `%s`? <%t>", tc.ErrMsg, expr, result)
				}
			}
		})
	}
}
func TestExecuteAllTestsInDir(t *testing.T) {
	dirs := &OrderedMap{
		om: map[int]string{
			1: "../test_folder/level2",
			2: "../test_folder/level3/sub-level",
			3: "../test_folder/level3/sub-level/sub-level2",
		},
		order: []int{1, 2, 3},
	}

	testCases := []struct {
		Index   int
		Name    string
		ErrMsg  string
		Results map[string]bool
	}{
		{
			Index: 1,
			Name:  "Level 2 Nesting",
			Results: map[string]bool{
				"TestFunc4":          false,
				"TestFile5":          false,
				"-- PASS: TestFunc3": true,
			},
			ErrMsg: "Level 2 nesting failed!",
		},
		{
			Index: 2,
			Name:  "Level 3.1 Nesting (Don't skip)",
			Results: map[string]bool{
				"-- FAIL: TestFile7": true,
			},
			ErrMsg: "Level 3.1 nesting failed!",
		},
		{
			Index: 3,
			Name:  "Level 3.2 Nesting",
			Results: map[string]bool{
				"TestFile7":          false,
				"-- PASS: TestFunc8": true,
			},
			ErrMsg: "Level 3.2 nesting failed!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual, _ := ExecuteAllTestsInDir(dirs, tc.Index)
			for expr, result := range tc.Results {
				if strings.Contains(actual, expr) != result {
					t.Errorf("%s Result contains `%s`? <%t>", tc.ErrMsg, expr, result)
				}
			}
		})
	}
}

func TestExecuteDefaultTestsInFile(t *testing.T) {
	dirs := &OrderedMap{
		om: map[int]string{
			1: "../test_folder/level1/file1_test.go",
			2: "../test_folder/level1/file2_test.go",
		},
		order: []int{1, 2},
	}

	testCases := []struct {
		Index int
		Name  string
		Expr  string
	}{
		{
			1,
			"Test should pass",
			"-- PASS: TestFunc1",
		},
		{
			2,
			"Test should be skipped",
			"-- SKIP: TestFile2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual, _ := ExecuteDefaultTestsInFile(dirs, tc.Index)
			fmt.Println(actual)
			if !strings.Contains(actual, tc.Expr) {
				t.Error("Not found!: " + tc.Expr)
			}
		})
	}
}

func TestExecuteAllTestsInFile(t *testing.T) {
	dirs := &OrderedMap{
		om: map[int]string{
			1: "../test_folder/level1/file1_test.go",
			2: "../test_folder/level1/file2_test.go",
		},
		order: []int{1, 2},
	}

	testCases := []struct {
		Index int
		Name  string
		Expr  string
	}{
		{
			1,
			"Test should pass",
			"-- PASS: TestFunc1",
		},
		{
			2,
			"Test should fail",
			"-- FAIL: TestFile2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual, _ := ExecuteAllTestsInFile(dirs, tc.Index)
			if !strings.Contains(actual, tc.Expr) {
				t.Error("Not found!: " + tc.Expr)
			}
		})
	}
}
