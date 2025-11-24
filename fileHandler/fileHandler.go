package fileHandler

import (
	"encoding/json"
	"fmt"
	"os"
	"taskTracker/task"
)

const fileName string = "taskTracker.json"

func New() {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		err = os.WriteFile(fileName, []byte("[]"), 0666)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func WriteFile(data any) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, b, 0666)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile() []task.Task {
	var task []task.Task
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &task)
	if err != nil {
		fmt.Println(err)
	}
	return task
}
