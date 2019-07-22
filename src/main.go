package main

import (
	"fmt"

	"github.com/dggr8/spanish-mem/src/cli"
	"github.com/dggr8/spanish-mem/src/file_operations"
	"github.com/dggr8/spanish-mem/src/language_test"
)

func main() {
	// Get all the data files.
	number_of_words := file_operations.LoadFiles()
	fmt.Println("I have", number_of_words, "words!")
	language_test.SeedWithTime()
	file_operations.GetWords()
	var command string
	for command != "exit" {
		switch command = cli.GetCommand(cli.Stdin, cli.Stdout); command {
		case "train spanish":
			language_test.TestSpanish()
		case "train english":
			language_test.TestEnglish()
		}
	}
}
