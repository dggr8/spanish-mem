package language

import (
	"fmt"
	"io"
	"time"

	"github.com/agnivade/levenshtein"
	"github.com/dggr8/spanish-mem/src/cli"
	"github.com/dggr8/spanish-mem/src/file_operations"
	"github.com/dggr8/spanish-mem/src/results"
)

func TestSpanish(rd io.Reader, wr io.Writer, resultJsonPath string) {
	trainCount := cli.GetInt(rd, wr, "How many words do you want to train")

	fmt.Fprintln(wr, "Translate these words to spanish:")
	TestLanguage(rd, wr, file_operations.EnglishToSpanish, trainCount, "spanish", resultJsonPath)
}

func TestEnglish(rd io.Reader, wr io.Writer, resultJsonPath string) {
	trainCount := cli.GetInt(rd, wr, "How many words do you want to train")

	fmt.Fprintln(wr, "Translate these words to english:")
	TestLanguage(rd, wr, file_operations.SpanishToEnglish, trainCount, "english", resultJsonPath)
}

func TestLanguage(rd io.Reader, wr io.Writer, LanguageMap map[string][]string,
	trainCount int, languageStr string, resultJsonPath string) {
	correctAnswers := 0
	LoopCtr := trainCount
	if trainCount <= 0 {
		return
	}

	for originalWord, translatedWords := range LanguageMap {
		answer := cli.GetAnswer(rd, wr, originalWord)
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
	_ = results.RecordResult(results.TestResult{
		Correct:   correctAnswers,
		Attempts:  trainCount,
		Train:     languageStr,
		Timestamp: time.Now(),
	}, resultJsonPath)
}

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
