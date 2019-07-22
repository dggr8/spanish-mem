package file_operations

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type WordPair struct {
	English string
	Spanish string
}

const globPath string = "../data/manually-curated/*.csv"

var FolderData = make([]WordPair, 0)
var SpanishToEnglish = make(map[string][]string)
var EnglishToSpanish = make(map[string][]string)

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

func formatFilename(file_path string) string {
	parts := strings.Split(file_path, "/")
	csvfile := parts[len(parts)-1]
	return csvfile[:len(csvfile)-4]
}

func GetWords() {
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
			EnglishWord := line[0]
			SpanishWord := line[1][1:]

			if SpanishToEnglish[SpanishWord] == nil {
				SpanishToEnglish[SpanishWord] = []string{EnglishWord}
			} else {
				SpanishToEnglish[SpanishWord] = append(SpanishToEnglish[SpanishWord], EnglishWord)
			}

			if EnglishToSpanish[EnglishWord] == nil {
				EnglishToSpanish[EnglishWord] = []string{SpanishWord}
			} else {
				EnglishToSpanish[EnglishWord] = append(EnglishToSpanish[EnglishWord], SpanishWord)
			}
		}
	}
}
