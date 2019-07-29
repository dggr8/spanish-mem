package language

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GlobPath is a link to the csv files.
const GlobPath string = "../data/manually-curated/*.csv"

// ParentPath is a link to dir that contains all the testing dirs.
const ParentPath string = "../data/"

// SpanishToEnglish and EnglishToSpanish maps are picked up by the testing packages.
var SpanishToEnglish = make(map[string][]string)
var EnglishToSpanish = make(map[string][]string)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// SwitchFolders allows users to switch between testing different directories.
func SwitchFolders(rd io.Reader, wr io.Writer, parentPath string) {

	allFileInfos, err := ioutil.ReadDir(parentPath)
	if err != nil {
		panic(err)
	}
	listOfDirs := make([]string, 0)
	for _, file := range allFileInfos {
		if file.IsDir() {
			listOfDirs = append(listOfDirs, file.Name())
		}
	}

	chosenDir := GetDirChoice(rd, wr, listOfDirs)
	GetWords(parentPath + "/" + chosenDir + "/*.csv")
}

// GetWords loads csv files from the globPath and populates SpanishToEnglish and
// EnglishToSpanish maps.
func GetWords(globPath string) {

	SpanishToEnglish = make(map[string][]string)
	EnglishToSpanish = make(map[string][]string)
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
