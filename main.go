package main

import (
	"flag"
	"fmt"
	"taskTracker/argsHandler"
	"taskTracker/fileHandler"
	"taskTracker/task"
)

func getTaskDetails() (int, []task.Task) {
	data := fileHandler.ReadFile()
	if len(data) != 0 {
		return data[len(data)-1].Id + 1, data
	} else {
		return 0, data
	}
}

func main() {
	fileHandler.New()
	id, tasks := getTaskDetails()
	flag.Usage = func() { fmt.Println(argsHandler.HelperText()) }
	flag.Parse()

	// validate arguments
	err := argsHandler.ValidateArgs()
	if err != nil {
		fmt.Println(err)
	} else {
		err := argsHandler.ExecuteCase(id, tasks)
		if err != nil {
			fmt.Println(err)
		}
	}

}
