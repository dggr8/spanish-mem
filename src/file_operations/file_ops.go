package file_operations

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

const GlobPath string = "../data/manually-curated/*.csv"

var SpanishToEnglish = make(map[string][]string)
var EnglishToSpanish = make(map[string][]string)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
