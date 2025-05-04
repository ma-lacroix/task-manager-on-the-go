package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"taskManager/data"
	"taskManager/userInput"
)

type REPL struct {
	taskManager data.TaskData
}

func NewReply() REPL {
	return REPL{*data.NewTaskData()}
}

func (r *REPL) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	menu := NewMenu()
	fmt.Println("Welcome to GoTaskManager. Type 'exit' to quit. \n")
	for {
		menu.ShowAllOptions()
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "exit" {
			break
		}
		menu.ShowOption(input)
		extraOptions := ""
		if input != "1" { // don't need additional input
			if !scanner.Scan() {
				break
			}
			extraOptions = scanner.Text()
			if extraOptions == "exit" {
				break
			}
		}
		userInputHandler := userInput.NewUserInput(input, strings.Split(extraOptions, " "))
		if !userInputHandler.HandleInput(&r.taskManager) {
			break
		}
	}
}
