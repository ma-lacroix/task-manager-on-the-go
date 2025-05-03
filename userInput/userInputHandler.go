package userInput

import (
	"fmt"
	"taskManager/tasks"
)

type UserInput struct {
	selection    Selection
	extraOptions []string
}

type Selection string

const (
	VIEW_ALL     Selection = "view_all"
	ADD_NEW      Selection = "add_new"
	DELETE_ENTRY Selection = "delete_entry"
)

func getSelection(s string) Selection {
	switch s {
	case "1":
		return VIEW_ALL
	case "2":
		return ADD_NEW
	case "3":
		return DELETE_ENTRY
	default:
		return VIEW_ALL
	}
}

func NewUserInput(s string, e []string) *UserInput {
	selection := getSelection(s)
	return &UserInput{selection, e}
}

func (u *UserInput) validateInput(inputString []string) bool {
	return true
}

func (u *UserInput) ViewAllTasks() bool {
	// TODO: show all tasks from DB here
	fmt.Println("All tasks here")
	return true
}

func (u *UserInput) AddNewTask() bool {
	task, err := tasks.NewTask(u.extraOptions[0], u.extraOptions[1])
	if err != nil {
		fmt.Println(err)
	}
	// TODO: write new task in DB
	fmt.Println("New task here ", task.Type, task.DueDate)
	return true
}

func (u *UserInput) DeleteEntryTask() bool {
	// TODO: remove data from DB
	fmt.Println("Deleting entry ", u.extraOptions[0])
	return true
}

func (u *UserInput) HandleInput() bool {
	status := true
	switch u.selection {
	case VIEW_ALL:
		status = u.ViewAllTasks()
	case ADD_NEW:
		status = u.AddNewTask()
	case DELETE_ENTRY:
		status = u.DeleteEntryTask()
	default:
		fmt.Println("Invalid operation, something went wrong")
		status = false
	}
	return status
}
