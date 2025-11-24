package argsHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"taskTracker/fileHandler"
	"taskTracker/task"
)

const (
	add            string = "add"
	update         string = "update"
	delete         string = "delete"
	markInProgress string = "mark-in-progress"
	markDone       string = "mark-done"
	list           string = "list"
)

func checkValidOption() bool {
	return os.Args[1] != add &&
		os.Args[1] != update &&
		os.Args[1] != delete &&
		os.Args[1] != markInProgress &&
		os.Args[1] != markDone &&
		os.Args[1] != list
}

func ValidateArgs() error {
	if len(os.Args) <= 1 {
		return errors.New("no arguments found")
	} else if os.Args[1] == add && len(os.Args) != 3 {
		err := fmt.Sprintf("Invalid argument usage %v. Usage - <taskTracker> %v <description>", add, add)
		return errors.New(err)
	} else if os.Args[1] == update && len(os.Args) != 4 {
		err := fmt.Sprintf("Invalid argument usage %v. Usage - <taskTracker> %v <id> <description>", update, update)
		return errors.New(err)
	} else if os.Args[1] == delete && len(os.Args) != 3 {
		err := fmt.Sprintf("Invalid argument usage %v. Usage - <taskTracker> %v <id>", delete, delete)
		return errors.New(err)
	} else if os.Args[1] == markInProgress && len(os.Args) != 3 {
		err := fmt.Sprintf("Invalid argument usage %v. Usage - <taskTracker> %v <id>", markInProgress, markInProgress)
		return errors.New(err)
	} else if os.Args[1] == markDone && len(os.Args) != 3 {
		err := fmt.Sprintf("Invalid argument usage %v. Usage - <taskTracker> %v <id>", markDone, markDone)
		return errors.New(err)
	} else if os.Args[1] == list && len(os.Args) >= 4 {
		err := fmt.Sprintf("Invalid argument usage %v. Usage - <taskTracker> %v or <taskTracker> %v <status> ", list, list, list)
		return errors.New(err)
	} else if checkValidOption() {
		err := fmt.Sprintf("Invalid argument %v", os.Args[1])
		return errors.New(err)
	} else {
		return nil
	}
}

func ExecuteCase(id int, tasks []task.Task) error {
	switch os.Args[1] {
	case add:
		task := task.Add(id, os.Args[2])
		err := fileHandler.WriteFile(append(tasks, task))
		if err != nil {
			return err
		}
	case update:
		id, _ := strconv.Atoi(os.Args[2])
		data, err := task.Update(id, tasks, os.Args[3])
		if err != nil {
			return err
		}
		err = fileHandler.WriteFile(data)
		if err != nil {
			return err
		}
	case delete:
		id, _ := strconv.Atoi(os.Args[2])
		data, err := task.Delete(id, tasks)
		if err != nil {
			return err
		}
		fileHandler.WriteFile(data)
	case markDone, markInProgress:
		id, _ := strconv.Atoi(os.Args[2])
		data, err := task.UpdateStatus(id, tasks, os.Args[1])
		if err != nil {
			return err
		}
		err = fileHandler.WriteFile(data)
		if err != nil {
			return err
		}
	case list:
		if len(os.Args) == 3 {
			tasks, err := task.ListByStatus(tasks, os.Args[2])
			if err != nil {
				return err
			} else {

				if len(tasks) == 0 {
					fmt.Println(string("[]"))
					return nil
				}
				data, err := json.MarshalIndent(tasks, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(data))
			}
		} else {
			data, err := json.MarshalIndent(tasks, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
		}

	default:
		fmt.Print("something went wrong")
	}
	return nil
}

func HelperText() string {
	text := `Welcome to Task Tracker
Task Tracker is a tool for managing daily tasks Tasks.
Usage:
        taskTracker <command>
The commands are:
    add                 Add a new Task
    update              Update any existing task
    delete              delete any task
    mark-in-progress    Mark in-progress status to a task
    mark-done           Mark done status to a task
    list                list all tasks`
	return text
}
