package main

import (
  "encoding/csv"
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  file, err := os.Open("../data/numbers.csv")
  check(err)
  defer file.Close()

  lines, err := csv.NewReader(file).ReadAll()
  check(err)
  fmt.Println("The number of entries are", len(lines))
  for _, line := range lines{
    fmt.Println(line[0] + " is" + line[1])
  }
}
