package main

import (
  "encoding/csv"
  "path/filepath"
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  // Get all the data files.
  all_files, err := filepath.Glob("../data/*.csv")
  check(err)
  for _, filename := range all_files {
    file, err := os.Open(filename)
    check(err)
    defer file.Close()

    lines, err := csv.NewReader(file).ReadAll()
    check(err)
    fmt.Println("The number of entries are", len(lines))
    for _, line := range lines{
      fmt.Println(line[0] + " is" + line[1])
    }
  }
}
