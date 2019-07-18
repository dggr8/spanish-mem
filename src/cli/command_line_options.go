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
	"list files":    true,
	"list words":    true,
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

func GetAnswer(question string) string {
	fmt.Print(question, "->")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func GetInt(question string) int {
	fmt.Print(question, "->")
	var number int
	fmt.Scanf("%d", &number)
	return number
}
