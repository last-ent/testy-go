package main

import (
	"os"

	"github.com/last-ent/testy-go/cli"
)

func main() {
	opts, err := cli.ParseFlags(os.Getwd)
	if err != nil {
		panic("Error while parsing flags: " + err.Error())
	}
	cli.StartPrompt(opts.Dir)
}
