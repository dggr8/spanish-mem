package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var list_of_commands = map[string]bool{
	"list files": true,
	"list words": true,
	"train spanish": true,
	"train english": true,
	"results": true,
	"exit": true,
}

func GetCommand() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What do you want to do now?")
	fmt.Print(">")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if list_of_commands[text] == true {
		return text
	} else {
		fmt.Println("I'm sorry. I didn't get that.")
		return ""
	}
}

func GetAnswer(question string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question, "->")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
