package language_test

import (
	"fmt"
	"github.com/agnivade/levenshtein"
  	"github.com/dggr8/spanish-mem/src/cli"
  	"github.com/dggr8/spanish-mem/src/file_operations"
)

func TestSpanish() {
	correct_count := 0
	word_list := file_operations.GetWords()
	fmt.Println("Translate these words to spanish:")
	for _, word_pair := range word_list {
		answer := cli.GetAnswer(word_pair.English)
		distance := levenshtein.ComputeDistance(answer, word_pair.Spanish)
		if 2 * float64(distance)/float64(len(answer) + len(word_pair.Spanish)) < 0.1 {
			fmt.Println("Correct!")
			correct_count = correct_count + 1
		} else {
			fmt.Printf("Nah! It is \"%v\".\n", word_pair.Spanish)
		}
	}
	fmt.Printf("%v correct out of %v", correct_count, len(word_list))
}