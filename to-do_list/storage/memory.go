package storage

import (
	"fmt"
	"todo-api/models"
)

var tasks []models.Task
var nextID int = 1

func GetAllTasks() []models.Task {
	return tasks
}

func AddTask(name string) models.Task {
	task := models.Task{
		ID:   nextID,
		Name: name,
		Done: false,
	}

	tasks = append(tasks, task)
	nextID++
	return task
}

func UpdateTask(id int, name string, done bool) (models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = name
			tasks[i].Done = done
			return tasks[i], nil
		}
	}

	return models.Task{}, fmt.Errorf("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}
