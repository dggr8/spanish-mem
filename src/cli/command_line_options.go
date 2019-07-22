package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	Stdin  = os.Stdin
	Stdout = os.Stdout
)

var list_of_commands = map[string]bool{
	"train spanish": true,
	"train english": true,
	"print results": true,
	"exit":          true,
}

func GetCommand(rd io.Reader, wr io.Writer) string {
	fmt.Fprintln(wr, "What do you want to do now?")
	fmt.Fprint(wr, ">")
	reader := bufio.NewReader(rd)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if list_of_commands[text] == true {
		return text
	} else {
		fmt.Fprintln(wr, "I'm sorry. I didn't get that.")
		return ""
	}
}

func GetAnswer(rd io.Reader, wr io.Writer, question string) string {
	fmt.Fprint(wr, question, "->")
	reader := bufio.NewReader(rd)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func GetInt(rd io.Reader, wr io.Writer, question string) (int, error) {
	fmt.Fprint(wr, question, "->")
	var number int
	_, err := fmt.Fscanf(rd, "%d", &number)
	if err != nil {
		return 0, err
	}
	return number, nil
}
