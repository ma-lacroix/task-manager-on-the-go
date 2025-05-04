package cli

import (
	"fmt"
)

type Menu struct{}

func NewMenu() Menu {
	return Menu{}
}

func (*Menu) ShowAllOptions() {
	fmt.Println("Please Select an Option: ")
	fmt.Println("1 - See All")
	fmt.Println("2 - Add New Task")
	fmt.Println("3 - Delete Task")
	fmt.Print("> ")
}

func (*Menu) ShowOption(selection string) bool {
	status := true
	switch selection {
	case "1":
		fmt.Println("here are all the tasks you've created: \n")
	case "2":
		fmt.Println("Add a new task by the type and execution date: \n")
	case "3":
		fmt.Println("Enter the ID of the task you want to delete: \n")
	default:
		fmt.Println("invalid option! \n")
		status = false
	}
	fmt.Print("> ")
	return status
}
