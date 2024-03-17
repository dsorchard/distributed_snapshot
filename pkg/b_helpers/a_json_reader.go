package helpers

import (
	"encoding/json"
	"os"
	models "snapshot/pkg/a_models"
)

const path = "/Users/arjunsunilkumar/GolandProjects/0sysdev_dec/snapshot/cmd/server/"

func readFile(fileName string) []byte {
	data, err := os.ReadFile(path + fileName)
	if err != nil {
		panic(err)
	}

	return data
}

func ReadNetConfig(fileName string) []models.ProcessInfo {
	data := readFile(fileName)

	var myJson []models.ProcessInfo
	err := json.Unmarshal(data, &myJson)
	if err != nil {
		panic(err)
	}

	return myJson
}

func ReadTaskList(fileName string) []models.Task {
	data := readFile(fileName)

	var myJson []models.Task
	err := json.Unmarshal(data, &myJson)
	if err != nil {
		panic(err)
	}

	return myJson
}
