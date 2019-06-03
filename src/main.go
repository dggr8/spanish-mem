package main

import (
  "fmt"
  "github.com/dggr8/spanish-mem/src/cli"
  "github.com/dggr8/spanish-mem/src/file_operations"
)

func main() {
  // Get all the data files.
  const file_glob string = "../data/*.csv"
  number_of_words := file_operations.LoadFiles(file_glob)
  fmt.Println("I have", number_of_words, "words!")
  var command string
  for command != "exit" {
    command = cli.GetCommand()
  }
}
