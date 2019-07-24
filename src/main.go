package main

import (
	"github.com/dggr8/spanish-mem/src/cli"
	"github.com/dggr8/spanish-mem/src/fileops"
	"github.com/dggr8/spanish-mem/src/language"
	"github.com/dggr8/spanish-mem/src/results"
)

func main() {
	// Get all the data files.
	fileops.GetWords(fileops.GlobPath)
	var command string
	for command != "exit" {
		switch command = cli.GetCommand(cli.Stdin, cli.Stdout); command {
		case "switch dirs":
			fileops.SwitchFolders(fileops.ParentPath)
		case "train spanish":
			language.TestSpanish(cli.Stdin, cli.Stdout, results.ResultJSONPath)
		case "train english":
			language.TestEnglish(cli.Stdin, cli.Stdout, results.ResultJSONPath)
		}
	}
}
