// Package cli is responsible for providing the UI through which the user can run his tests.
package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/last-ent/testy-go/classifier"
	"github.com/last-ent/testy-go/commands"

	prompt "github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	for _, s := range suggestionsMap {
		suggestions = append(suggestions, s...)
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func getExecutor(dir string) func(string) {
	return func(s string) {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		dirMeta, err := classifier.TraverseDir(dir)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse provided directory. Dir: %s\nError: %s\n", dir, err))
		}

		switch s {
		case "":
			return
		case "quit", "exit", "q":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			result, err := commandProcessor(dirMeta, strings.Fields(s))
			if err != nil {
				fmt.Printf("%s\n%s\n", result, err)
			} else {
				fmt.Println(result)
			}
		}
	}
}

func commandProcessor(dirMeta *classifier.TestableDirMeta, cmds []string) (string, error) {
	switch cmds[0] {
	case "list":
		return commands.ListAll(dirMeta), nil
	case "list-dirs":
		return commands.ListDirs(dirMeta), nil
	case "list-files":
		return commands.ListFiles(dirMeta), nil
	case "test":
		return commands.ExecuteDefaultTests(dirMeta.ParentDir)
	case "test-all":
		return commands.ExecuteAllTests(dirMeta.ParentDir)
	case "test-dir":
		return executeDirTests(cmds, dirMeta, commands.ExecuteDefaultTestsInNestedDir)
	case "test-dir-all":
		return executeDirTests(cmds, dirMeta, commands.ExecuteAllTestsInNestedDir)
	case "test-only-dir":
		return executeDirTests(cmds, dirMeta, commands.ExecuteDefaultTestsInDir)
	case "test-only-dir-all":
		return executeDirTests(cmds, dirMeta, commands.ExecuteAllTestsInDir)
	case "test-file":
		return executeFileTests(cmds, dirMeta, commands.ExecuteDefaultTestsInFile)
	case "test-file-all":
		return executeFileTests(cmds, dirMeta, commands.ExecuteAllTestsInFile)
	default:
		return "", fmt.Errorf("Unknown command option")
	}
}

func executeDirTests(cmds []string, dirMeta *classifier.TestableDirMeta, executionFunction func(map[int]string, int) (string, error)) (string, error) {
	index, err := getIndexNumber(cmds)
	if err != nil {
		return "", err
	}
	mp := commands.ListDirsAsMap(dirMeta)
	return executionFunction(mp, index)
}

func executeFileTests(cmds []string, dirMeta *classifier.TestableDirMeta, executionFunction func(map[int]string, int) (string, error)) (string, error) {
	index, err := getIndexNumber(cmds)
	if err != nil {
		return "", err
	}
	mp := commands.ListFilesAsMap(dirMeta)
	return executionFunction(mp, index)
}

func getIndexNumber(cmds []string) (int, error) {
	switch len(cmds) {
	case 0, 1:
		return -1, fmt.Errorf("please provide an index number")
	default:
		i := cmds[1]
		index, err := strconv.Atoi(i)
		if err != nil {
			return -1, fmt.Errorf("please provide a valid index number")
		}
		return index, nil
	}
}

// StartPrompt starts the repl prompt from which the application starts.
func StartPrompt(dir string) {
	p := prompt.New(
		getExecutor(dir),
		completer,
		prompt.OptionTitle("Welcome to Testy Go! A tool for easy testing."),
		prompt.OptionPrefix(">>> "),
	)
	p.Run()
}
