package language_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/agnivade/levenshtein"
	"github.com/dggr8/spanish-mem/src/cli"
	"github.com/dggr8/spanish-mem/src/file_operations"
	"github.com/dggr8/spanish-mem/src/results"
)

func SeedWithTime() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestSpanish() {
	correctAnswers := 0
	trainCount := cli.GetInt(
		cli.Stdin, cli.Stdout, "How many words do you want to train")
	LoopCtr := trainCount

	fmt.Println("Translate these words to spanish:")
	for englishWord, SpanishList := range file_operations.EnglishToSpanish {
		answer := cli.GetAnswer(cli.Stdin, cli.Stdout, englishWord)
		if answer == "" {
			fmt.Printf("It is one of \"%v\".\n", SpanishList)
			continue
		}
		distance := MinDistance(answer, SpanishList)
		if distance == 0 {
			fmt.Println("Correct!")
			correctAnswers = correctAnswers + 1
		} else {
			fmt.Printf("Nah! It is one of \"%v\".\n", SpanishList)
		}

		LoopCtr--
		if LoopCtr == 0 {
			break
		}
	}
	fmt.Printf("%v correct out of %v\n", correctAnswers, trainCount)
	_ = results.RecordResult(results.TestResult{
		Correct:   correctAnswers,
		Attempts:  trainCount,
		Train:     "spanish",
		Timestamp: time.Now(),
	}, results.ResultFilePath)
}

func TestEnglish() {
	correctAnswers := 0
	trainCount := cli.GetInt(
		cli.Stdin, cli.Stdout, "How many words do you want to train")
	LoopCtr := trainCount

	fmt.Println("Translate these words to english:")
	for spanishWord, EnglishList := range file_operations.SpanishToEnglish {
		answer := cli.GetAnswer(cli.Stdin, cli.Stdout, spanishWord)
		if answer == "" {
			fmt.Printf("It is one of \"%v\".\n", EnglishList)
			continue
		}
		distance := MinDistance(answer, EnglishList)
		if distance == 0 {
			fmt.Println("Correct!")
			correctAnswers = correctAnswers + 1
		} else {
			fmt.Printf("Nah! It is one of \"%v\".\n", EnglishList)
		}

		LoopCtr--
		if LoopCtr == 0 {
			break
		}
	}
	fmt.Printf("%v correct out of %v\n", correctAnswers, trainCount)
	_ = results.RecordResult(results.TestResult{
		Correct:   correctAnswers,
		Attempts:  trainCount,
		Train:     "english",
		Timestamp: time.Now(),
	}, results.ResultFilePath)
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
