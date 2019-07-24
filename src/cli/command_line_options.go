package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Stdin and Stdout are global variables to access os counterparts.
var (
	Stdin  = os.Stdin
	Stdout = os.Stdout
)

var listOfCommands = map[string]bool{
	"train spanish": true,
	"train english": true,
	"print results": true,
	"exit":          true,
}

// GetCommand asks the user what they want to do and returns the command
// as a string. If the command isn't in listOfCommands, the user is prompted again.
func GetCommand(rd io.Reader, wr io.Writer) string {
	fmt.Fprintln(wr, "What do you want to do now?")
	fmt.Fprint(wr, ">")
	reader := bufio.NewReader(rd)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if listOfCommands[text] == true {
		return text
	} else {
		fmt.Fprintln(wr, "I'm sorry. I didn't get that.")
		return ""
	}
}

// GetAnswer asks the user for the answer to an input question and returns the answer
// as a string.
func GetAnswer(rd io.Reader, wr io.Writer, question string) string {
	fmt.Fprint(wr, question, "->")
	reader := bufio.NewReader(rd)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

// GetInt asks the user an input question and returns an integer.
func GetInt(rd io.Reader, wr io.Writer, question string) (number int) {
	for {
		fmt.Fprint(wr, question, "->")
		_, err := fmt.Fscanf(rd, "%d", &number)
		if err == nil && number >= 0 {
			return
		}
	}
}
