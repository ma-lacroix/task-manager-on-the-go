package data

import (
	"fmt"
	"strconv"
	"taskManager/tasks"
)

type TaskData struct {
	taskMap map[int]tasks.Task
	counter int
}

func NewTaskData() *TaskData {
	return &TaskData{make(map[int]tasks.Task), 0}
}

func (t *TaskData) AddTask(newTaskToAdd tasks.Task) bool {
	t.taskMap[t.counter] = newTaskToAdd
	t.counter++
	return true
}

func (t *TaskData) GetTaskMap() bool {
	for k, v := range t.taskMap {
		fmt.Println(k, v)
	}
	return true
}

func (t *TaskData) RemoveTask(key string) bool {
	i, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
		return false
	}
	delete(t.taskMap, i)
	return true
}
