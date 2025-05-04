package userInput

import (
	"fmt"
	"taskManager/data"
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

func (u *UserInput) ViewAllTasks(taskData *data.TaskData) bool {
	fmt.Println("All tasks here \n")
	return taskData.GetTaskMap()
}

func (u *UserInput) AddNewTask(taskData *data.TaskData) bool {
	task, err := tasks.NewTask(u.extraOptions[0], u.extraOptions[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("New task added \n", task.Type, task.DueDate)
	return taskData.AddTask(*task)
}

func (u *UserInput) DeleteEntryTask(taskData *data.TaskData) bool {
	fmt.Println("Deleting entry \n", u.extraOptions[0])
	return taskData.RemoveTask(u.extraOptions[0])
}

func (u *UserInput) HandleInput(taskData *data.TaskData) bool {
	status := true
	switch u.selection {
	case VIEW_ALL:
		status = u.ViewAllTasks(taskData)
	case ADD_NEW:
		status = u.AddNewTask(taskData)
	case DELETE_ENTRY:
		status = u.DeleteEntryTask(taskData)
	default:
		fmt.Println("Invalid operation, something went wrong \n")
		status = false
	}
	return status
}
