package cli

import (
	prompt "github.com/c-bata/go-prompt"
)

type suggestionType int

const (
	// List is key for List suggestions
	List suggestionType = iota

	// Test is key for Test suggestions
	Test

	// Others is key for various arbitrary suggestions
	Others
)

var listSuggestions = []prompt.Suggest{
	{Text: "list", Description: "List all packages and source files that can be tested."},
	{Text: "list-dirs", Description: "List all packages that can be tested."},
	{Text: "list-files", Description: "List all source files that can be tested."},
}

var testSuggestions = []prompt.Suggest{
	{Text: "test", Description: "Run all non-skippable tests under `dir`."},
	{Text: "test-all", Description: "Run all tests (including skippable) under `dir`."},
	{Text: "test-dir", Description: "Run all non-skippable tests under package (selected by index number) including nested sub-packages."},
	{Text: "test-dir-all", Description: "Run all tests (including skippable) under package (selected by index number) including nested sub-packages."},
	{Text: "test-only-dir", Description: "Run all tests immediately under package (selected by index number)"},
	{Text: "test-only-dir-all", Description: "Run all tests (including skippable) immediately under package (selected by index number)"},
	{Text: "test-file", Description: "Run all non-skippable tests in the file selected. (Selected by index number)"},
	{Text: "test-file-all", Description: "Run all tests (including skippable) in the file selected. (Selected by index number)"},
}

var otherSuggestions = []prompt.Suggest{
	{Text: "q", Description: "Quit/Exit application"},
	{Text: "quit", Description: "Quit/Exit application"},
	{Text: "exit", Description: "Quit/Exit application"},
}

var suggestionsMap = map[suggestionType][]prompt.Suggest{
	List:   listSuggestions,
	Test:   testSuggestions,
	Others: otherSuggestions,
}
