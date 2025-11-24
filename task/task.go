package task

import (
	"errors"
	"strings"
	"time"
)

const (
	todo       string = "todo"
	inProgress string = "in-progress"
	done       string = "done"
)

type Task struct {
	Id          int       `json:"Id"`
	Description string    `json:"Description"`
	Status      string    `json:"Status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

func Add(id int, description string) Task {
	return Task{
		Id:          id,
		Description: description,
		Status:      todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func Delete(id int, tasks []Task) ([]Task, error) {
	for idx, val := range tasks {
		if val.Id == id {
			updatedTask := append(tasks[:idx], tasks[idx+1:]...)
			return updatedTask, nil
		}
	}
	return nil, errors.New("id not found")
}

func Update(id int, tasks []Task, description string) ([]Task, error) {
	for idx, val := range tasks {
		if val.Id == id {
			tasks[idx].Description = description
			tasks[idx].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return nil, errors.New("id not found")
}

func UpdateStatus(id int, tasks []Task, status string) ([]Task, error) {
	for idx, val := range tasks {
		if val.Id == id {
			tasks[idx].Status = strings.ReplaceAll(status, "mark-", "")
			tasks[idx].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return nil, errors.New("id not found")
}

func ListByStatus(tasks []Task, status string) ([]Task, error) {
	var filteredList []Task
	if status != inProgress && status != done && status != todo {
		return nil, errors.New("status argument must be either todo, in-progress or done")
	}
	for idx, val := range tasks {
		if val.Status == status {
			filteredList = append(filteredList[:], tasks[idx])
		}
	}
	return filteredList, nil
}
