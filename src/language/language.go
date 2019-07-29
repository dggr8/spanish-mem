// Package language provides functions to test between languages.
package language

import (
	"fmt"
	"io"
	"time"

	"github.com/agnivade/levenshtein"
)

// TestSpanish tests the user's ability to translate from English to Spanish.
// It takes in io.Reader, io.Writer and the path to the results.json file.
func TestSpanish(rd io.Reader, wr io.Writer, resultJSONPath string) {
	trainCount := GetInt(rd, wr, "How many words do you want to train")

	fmt.Fprintln(wr, "Translate these words to spanish:")
	TestLanguage(rd, wr, EnglishToSpanish, trainCount, "spanish", resultJSONPath)
}

// TestEnglish tests the user's ability to translate Spanish to English.
// It takes in io.Reader, io.Writer and the path to the results.json file.
func TestEnglish(rd io.Reader, wr io.Writer, resultJSONPath string) {
	trainCount := GetInt(rd, wr, "How many words do you want to train")

	fmt.Fprintln(wr, "Translate these words to english:")
	TestLanguage(rd, wr, SpanishToEnglish, trainCount, "english", resultJSONPath)
}

// TestLanguage provides a generic function to test conversion between two languages.
// It takes in a map[string][]string as database to test on, number of conversions to test and
// string tag to be added into the result.
func TestLanguage(rd io.Reader, wr io.Writer, LanguageMap map[string][]string,
	trainCount int, languageStr string, resultJSONPath string) {
	correctAnswers := 0
	LoopCtr := trainCount
	if trainCount <= 0 {
		return
	}

	for originalWord, translatedWords := range LanguageMap {
		answer := GetAnswer(rd, wr, originalWord)
		if answer == "" {
			fmt.Fprintf(wr, "It is one of \"%v\".\n", translatedWords)
			continue
		}
		distance := MinDistance(answer, translatedWords)
		if distance == 0 {
			fmt.Fprintln(wr, "Correct!")
			correctAnswers++
		} else {
			fmt.Fprintf(wr, "Nah! It is one of \"%v\".\n", translatedWords)
		}

		LoopCtr--
		if LoopCtr == 0 {
			break
		}
	}

	fmt.Fprintf(wr, "%v correct out of %v\n", correctAnswers, trainCount)
	_ = RecordResult(TestResult{
		Correct:   correctAnswers,
		Attempts:  trainCount,
		Train:     languageStr,
		Timestamp: time.Now(),
	}, resultJSONPath)
}

// MinDistance finds the least levenshtein distance between the word and all the
// words in comparingWords.
func MinDistance(word string, comparingWords []string) (min int) {
	min = len(word)
	for _, compareWord := range comparingWords {
		distance := levenshtein.ComputeDistance(word, compareWord)
		if min > distance {
			min = distance
		}
	}
	return
}
