// Package file_operations loads the maps of translations.
package fileops

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

// GlobPath is a link to the csv files.
const GlobPath string = "../data/manually-curated/*.csv"

// SpanishToEnglish and EnglishToSpanish maps are picked up by the testing packages.
var SpanishToEnglish = make(map[string][]string)
var EnglishToSpanish = make(map[string][]string)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetWords loads csv files from the globPath and populates SpanishToEnglish and
// EnglishToSpanish maps.
func GetWords(globPath string) {
	allFiles, err := filepath.Glob(globPath)
	check(err)
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
