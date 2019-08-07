package language

// cli provides functions to ask for and recieve user inputs.
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/nsf/termbox-go"
)

// Stdin and Stdout are global variables to access os counterparts.
var (
	Stdin  = os.Stdin
	Stdout = os.Stdout
)

var listOfCommands = map[string]bool{
	"switch dirs":   true,
	"train spanish": true,
	"train english": true,
	"print results": true,
	"exit":          true,
}

var au aurora.Aurora

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
	}
	fmt.Fprintln(wr, "I'm sorry. I didn't get that.")
	return ""
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

// GetDirChoice offers the user a choice of dir to train on. Returns a string
// of user's choice.
func GetDirChoice(rd io.Reader, wr io.Writer, listOfDir []string) string {
	if len(listOfDir) == 1 {
		fmt.Fprintf(wr, "No choice in directories.\n")
		return listOfDir[0]
	}

	pos := 0
	for {
		termbox.Init()
		err := printList(wr, listOfDir, pos)
		if err != nil {
			return listOfDir[pos]
		}
		fmt.Fprintln(wr, "Select one of the directories (highlighted in red) and hit enter!")
		event := termbox.PollEvent()
		if event.Key == termbox.KeyEnter {
			termbox.Close()
			return listOfDir[pos]
		} else if event.Key == termbox.KeyArrowUp || event.Ch == 'k' {
			if pos > 0 {
				pos--
			}
		} else if event.Key == termbox.KeyArrowDown || event.Ch == 'j' {
			if pos < len(listOfDir[pos])-1 {
				pos++
			}
		}
		termbox.Close()
	}
}

func printList(wr io.Writer, listOfString []string, pos int) error {
	if pos < 0 || pos >= len(listOfString) {
		return errors.New("pos is out of range")
	}

	au = aurora.NewAurora(true)
	for n, str := range listOfString {
		if n == pos {
			fmt.Fprintln(wr, au.Red(str))
		} else {
			fmt.Fprintln(wr, au.Cyan(str))
		}
	}
	return nil
}
