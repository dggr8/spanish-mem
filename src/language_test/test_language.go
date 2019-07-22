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
	correct_count := 0
	train_count, err := cli.GetInt(
		cli.Stdin, cli.Stdout, "How many words do you want to train")
	for err != nil || train_count <= 0 {
		train_count, err = cli.GetInt(
			cli.Stdin, cli.Stdout, "Oops.How many words so you want to train")
	}
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
			correct_count = correct_count + 1
		} else {
			fmt.Printf("Nah! It is %q.", SpanishList[0])
			fmt.Printf("The distance was %d.", distance)
		}

		train_count--
		if train_count == 0 {
			break
		}
	}
	fmt.Printf("%v correct out of %v\n", correct_count, train_count)
	results.RecordResult(correct_count, train_count, "spanish")
}

func TestEnglish() {
	correct_count := 0
	train_count, err := cli.GetInt(
		cli.Stdin, cli.Stdout, "How many words so you want to train")
	for err != nil {
		train_count, err = cli.GetInt(
			cli.Stdin, cli.Stdout, "Oops.How many words so you want to train")
	}
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
			correct_count = correct_count + 1
		} else {
			fmt.Printf("Nah! It is one of \"%v\".\n", EnglishList)
			fmt.Printf("The distance was %d.", distance)
		}

		train_count--
		if train_count == 0 {
			break
		}
	}
	fmt.Printf("%v correct out of %v\n", correct_count, train_count)
	results.RecordResult(correct_count, train_count, "english")
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
