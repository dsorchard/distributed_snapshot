package process

import (
	"fmt"
	models "snapshot/pkg/a_models"
	helpers "snapshot/pkg/b_helpers"
)

type Process struct {
	mainJob       *MainJob
	communication *CommunicationModule
	stateManager  *StateManager
}

func CreateProcess(processId int, network []models.ProcessInfo, taskList []models.Task, quit chan bool) *Process {
	processInfo := network[processId]
	fmt.Println(fmt.Sprintf("Creating process: %v, in port: %v", processInfo.Name, processInfo.Port))

	processMessageIn := make(chan models.Message)
	processMessageOut := make(chan models.Message)

	markMessageIn := make(chan models.Message)
	markMessageOut := make(chan int)

	updateStateChan := make(chan models.ProcessEvent)
	saveGlobalState := make(chan int)
	logger := helpers.CreateLogger(processInfo.Name)

	thisJob := CreateJob(processInfo, network, updateStateChan, processMessageIn, processMessageOut, taskList, saveGlobalState, quit)

	thisCommunicationMod := CreateCommunicationModule(processId, network, processMessageIn, processMessageOut, markMessageIn, markMessageOut, logger)

	thisStateManager := CreateStateManager(processId, updateStateChan, saveGlobalState, markMessageIn, markMessageOut, logger)

	thisProcess := Process{
		mainJob:       thisJob,
		communication: thisCommunicationMod,
		stateManager:  thisStateManager,
	}

	return &thisProcess
}
