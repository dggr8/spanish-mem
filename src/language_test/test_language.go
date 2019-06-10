package language_test

import (
	"fmt"
	"github.com/agnivade/levenshtein"
  	"github.com/dggr8/spanish-mem/src/cli"
	"github.com/dggr8/spanish-mem/src/file_operations"
  	"github.com/dggr8/spanish-mem/src/results"
	"math/rand"
	"time"
)

func SeedWithTime() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestSpanish() {
	correct_count := 0
	word_list := file_operations.GetWords()
	SeedWithTime()
	train_count := cli.GetInt("How many words do you want to train")
	fmt.Println("Translate these words to spanish:")
	for i := 0; i < train_count; i++ {
		word_pair := word_list[rand.Intn(len(word_list))]
		answer := cli.GetAnswer(word_pair.English)
		distance := levenshtein.ComputeDistance(answer, word_pair.Spanish)
		if 2 * float64(distance)/float64(len(answer) + len(word_pair.Spanish)) < 0.1 {
			fmt.Println("Correct!")
			correct_count = correct_count + 1
		} else {
			fmt.Printf("Nah! It is \"%v\".\n", word_pair.Spanish)
		}
	}
	fmt.Printf("%v correct out of %v\n", correct_count, train_count)
	results.RecordResult(correct_count, train_count, "spanish")
}

func TestEnglish() {
	correct_count := 0
	word_list := file_operations.GetWords()
	SeedWithTime()
	train_count := cli.GetInt("How many words so you want to train")
	fmt.Println("Translate these words to english:")
	for i := 0; i < train_count; i++ {
		word_pair := word_list[rand.Intn(len(word_list))]
		answer := cli.GetAnswer(word_pair.Spanish)
		distance := levenshtein.ComputeDistance(answer, word_pair.English)
		if 2 * float64(distance)/float64(len(answer) + len(word_pair.English)) < 0.1 {
			fmt.Println("Correct!")
			correct_count = correct_count + 1
		} else {
			fmt.Printf("Nah! It is \"%v\".\n", word_pair.English)
		}
	}
	fmt.Printf("%v correct out of %v\n", correct_count, train_count)
	results.RecordResult(correct_count, train_count, "english")
}