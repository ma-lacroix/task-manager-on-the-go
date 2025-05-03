package tasks

import (
	"errors"
	"strings"
)

type TaskType string

const (
	REMINDER TaskType = "REMINDER"
	ACTION   TaskType = "ACTION"
	DEADLINE TaskType = "DEADLINE"
)

type Task struct {
	Type    TaskType
	DueDate string
}

func getTaskType(t string) (TaskType, error) {
	switch strings.ToUpper(t) {
	case "REMINDER":
		return REMINDER, nil
	case "ACTION":
		return ACTION, nil
	case "DEADLINE":
		return DEADLINE, nil
	}
	return "", errors.New("invalid task type")
}

func NewTask(t string, dueDate string) (*Task, error) {
	taskType, err := getTaskType(t)
	if err != nil {
		return nil, errors.New("task type invalid")
	}
	return &Task{Type: taskType, DueDate: dueDate}, nil
}
