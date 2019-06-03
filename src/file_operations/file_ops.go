package file_operations

import (
	"encoding/csv"
	"fmt"
  "path/filepath"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
	  panic(e)
	}
}


func LoadFiles(glob_path string) int64 {
	// Get all the data files from the glob path.
	all_files, err := filepath.Glob(glob_path)
	check(err)
	var number_of_words int64
	for _, filename := range all_files {
		file, err := os.Open(filename)
		check(err)
		defer file.Close()

		lines, err := csv.NewReader(file).ReadAll()
		check(err)
		number_of_words += int64(len(lines))
	}

	return number_of_words
}

func ListFiles(glob_path string) {
	// Print a list of files that tests use.
	all_files, err := filepath.Glob(glob_path)
	check(err)
	fmt.Println("------------------------------------")
	for _, filepath := range all_files {
		fmt.Println(formatFilename(filepath))
	}
	fmt.Println("------------------------------------")
}

func formatFilename(file_path string) string {
	parts := strings.Split(file_path, "/")
	csvfile := parts[len(parts)-1]
	return csvfile[:len(csvfile)-4]
}

func ListWords(glob_path string) {
	// Print a list of words that tests use.
	all_files, err := filepath.Glob(glob_path)
	check(err)
	fmt.Println("------------------------------------")
	count := 1
	for _, filepath := range all_files {
		file, err := os.Open(filepath)
		check(err)
		defer file.Close()

		lines, err := csv.NewReader(file).ReadAll()
		check(err)
		for _, line := range lines {
			fmt.Println(count, line[0], "==", line[1])
			count = count + 1
		}
	}
	fmt.Println("------------------------------------")
}