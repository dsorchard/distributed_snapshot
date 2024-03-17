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

	// local timer task (just incr. amount) + processMessageIn task
	updateStateChan := make(chan models.ProcessEvent)

	// network task (appTask, snapTask)
	processMessageIn := make(chan models.Message)
	processMessageOut := make(chan models.Message)
	markMessageIn := make(chan models.Message)
	saveGlobalState := make(chan int)
	markMessageOut := make(chan int)

	logger := helpers.CreateLogger(processInfo.Name)

	thisCommunicationMod := CreateCommunicationModule(processId, network, processMessageIn, processMessageOut, markMessageIn, markMessageOut, logger)

	thisJob := CreateJob(processInfo, network, updateStateChan, processMessageIn, processMessageOut, taskList, saveGlobalState, quit)

	thisStateManager := CreateStateManager(processId, updateStateChan, saveGlobalState, markMessageIn, markMessageOut, logger)

	thisProcess := Process{
		mainJob:       thisJob,
		communication: thisCommunicationMod,
		stateManager:  thisStateManager,
	}

	return &thisProcess
}
