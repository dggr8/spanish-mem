package main

import (
	"github.com/dggr8/spanish-mem/src/language"
)

func main() {
	// Get all the data files.
	language.GetWords(language.GlobPath)
	var command string
	for command != "exit" {
		switch command = language.GetCommand(language.Stdin, language.Stdout); command {
		case "switch dirs":
			language.SwitchFolders(language.Stdin, language.Stdout, language.ParentPath)
		case "train spanish":
			language.TestSpanish(language.Stdin, language.Stdout, language.ResultJSONPath)
		case "train english":
			language.TestEnglish(language.Stdin, language.Stdout, language.ResultJSONPath)
		case "print results":
			language.PrintResults(language.Stdout, language.ResultJSONPath)
		}
	}
}
