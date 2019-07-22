package file_operations

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const globPath string = "../data/manually-curated/*.csv"

type WordPair struct {
	English string
	Spanish string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFiles() int64 {
	// Get all the data files from the glob path.
	allFiles, err := filepath.Glob(globPath)
	check(err)
	var number_of_words int64
	for _, filename := range allFiles {
		file, err := os.Open(filename)
		check(err)
		defer file.Close()

		lines, err := csv.NewReader(file).ReadAll()
		check(err)
		number_of_words += int64(len(lines))
	}

	return number_of_words
}

func ListFiles() {
	// Print a list of files that tests use.
	allFiles, err := filepath.Glob(globPath)
	check(err)
	fmt.Println("------------------------------------")
	for _, filepath := range allFiles {
		fmt.Println(formatFilename(filepath))
	}
	fmt.Println("------------------------------------")
}

func formatFilename(file_path string) string {
	parts := strings.Split(file_path, "/")
	csvfile := parts[len(parts)-1]
	return csvfile[:len(csvfile)-4]
}

func ListWords() {
	// Print a list of words that tests use.
	allFiles, err := filepath.Glob(globPath)
	check(err)
	fmt.Println("------------------------------------")
	count := 1
	for _, filepath := range allFiles {
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

func GetWords() []WordPair {
	var word_list []WordPair
	allFiles, err := filepath.Glob(globPath)
	check(err)
	fmt.Println("------------------------------------")
	for _, filepath := range allFiles {
		file, err := os.Open(filepath)
		check(err)
		defer file.Close()

		lines, err := csv.NewReader(file).ReadAll()
		check(err)
		for _, line := range lines {
			word_list = append(word_list, WordPair{line[0], line[1][1:]})
		}
	}
	return word_list
}
