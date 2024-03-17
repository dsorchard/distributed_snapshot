package main

import (
	"os"
	helpers "snapshot/pkg/b_helpers"
	process "snapshot/pkg/c_process"

	"strconv"
)

func main() {
	// Obtain process data
	args := os.Args[1:]
	processId := args[0]
	fileName := args[1]

	index, err := strconv.Atoi(processId)
	if err != nil {
		panic("Invalid argument when creating process")
	}

	network := helpers.ReadNetConfig(fileName)
	taskList := helpers.ReadTaskList("tasks/P" + processId + "Tasks.json")
	quit := make(chan bool)

	process.CreateProcess(index, network, taskList, quit)

	<-quit
}
