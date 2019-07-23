package main

import (
	"github.com/dggr8/spanish-mem/src/cli"
	"github.com/dggr8/spanish-mem/src/file_operations"
	"github.com/dggr8/spanish-mem/src/language"
	"github.com/dggr8/spanish-mem/src/results"
)

func main() {
	// Get all the data files.
	file_operations.GetWords(file_operations.GlobPath)
	var command string
	for command != "exit" {
		switch command = cli.GetCommand(cli.Stdin, cli.Stdout); command {
		case "train spanish":
			language.TestSpanish(cli.Stdin, cli.Stdout, results.ResultJsonPath)
		case "train english":
			language.TestEnglish(cli.Stdin, cli.Stdout, results.ResultJsonPath)
		}
	}
}
