package printer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	successPrinter = color.New(color.FgGreen)
	failPrinter    = color.New(color.FgHiRed)
	skipPrinter    = color.New(color.FgYellow)

	unsurePrinter = color.New(color.FgHiMagenta)
)

type jsonOutput struct {
	Time    string
	Action  string
	Package string
	Test    string
	Output  string
	Elapsed float64
}

type collectionItem struct {
	Outputs []string
	Action  string
}

type separatedCollections struct {
	Success map[string]collectionItem
	Fail    map[string]collectionItem
	Skip    map[string]collectionItem
}

func parseLine(collection map[string]collectionItem, line string) {
	var output jsonOutput
	if err := json.Unmarshal([]byte(line), &output); err != nil {
		OnFail(fmt.Sprintf("Unable to parse: %s.\nError: %s\n", line, err))
		return
	}

	testName := fmt.Sprintf("%s.%s", output.Package, output.Test)
	item := collection[testName]

	switch action := output.Action; action {
	case "pass", "fail", "skip":
		item.Action = action
	default:
		if item.Outputs == nil {
			item.Outputs = []string{}
		}
		item.Outputs = append(item.Outputs, output.Output)
	}

	collection[testName] = item
}

func separateItems(collection map[string]collectionItem) *separatedCollections {
	success := map[string]collectionItem{}
	fail := map[string]collectionItem{}
	skip := map[string]collectionItem{}

	for key, item := range collection {
		switch item.Action {
		case "pass":
			success[key] = item
		case "fail":
			fail[key] = item
		case "skip":
			skip[key] = item
		}
	}

	return &separatedCollections{
		Success: success,
		Fail:    fail,
		Skip:    skip,
	}
}

func printNonPass(collection map[string]collectionItem, printer func(string)) {
	for key, item := range collection {
		printer(fmt.Sprintf("%s\n%s\n", key, strings.Join(item.Outputs, "")))
	}
}

// PrintResults is used to print results after execution
func PrintResults(blob string) {
	lines := strings.Split(blob, "\n")

	plain := "\n"
	collection := map[string]collectionItem{}

	for _, line := range lines {
		if strings.HasPrefix(line, `{"Time"`) {
			parseLine(collection, line)
		} else {
			plain += line + "\n"
		}
	}

	printCollections(separateItems(collection))
	unsurePrinter.Print(plain)
}

func printCollections(collections *separatedCollections) {

	fmt.Println("\n============================")
	successPrinter.Println("SUCCESSFUL TESTS")
	fmt.Println("----------------------------")
	for testName := range collections.Success {
		successPrinter.Printf("%s\n", testName)
	}
	fmt.Println("============================")

	OnFail("FAILED TESTS")
	fmt.Println("----------------------------")
	printNonPass(collections.Fail, OnFail)
	fmt.Println("============================")

	OnSkip("SKIPPED TESTS")
	fmt.Println("----------------------------")
	printNonPass(collections.Skip, OnSkip)
	fmt.Println("============================")

}

// OnFail prints message with red coloring
func OnFail(msg string) {
	failPrinter.Println(msg)
}

// OnSkip prints message with yellow coloring
func OnSkip(msg string) {
	skipPrinter.Println(msg)
}
