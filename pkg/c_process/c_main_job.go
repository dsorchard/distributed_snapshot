package process

import (
	"math/rand"
	models "snapshot/pkg/a_models"
	"sync"
	"time"
)

// This element represents the main task of the Process
type MainJob struct {
	ProcessInfo       models.ProcessInfo   // it's own info
	NetworkInfo       []models.ProcessInfo // other processes info
	Data              models.ProcessData
	updateStateChan   chan models.ProcessEvent
	processMessageOut chan models.Message
	processMessageIn  chan models.Message
	saveGlobalState   chan int
	ticker            *time.Ticker
}

// Returns the pointer to a new MainJob struct
func CreateJob(processInfo models.ProcessInfo, network []models.ProcessInfo, updateStateChan chan models.ProcessEvent, processMessageIn chan models.Message, processMessageOut chan models.Message, taskList []models.Task, saveGlobalState chan int, quit chan bool) *MainJob {
	seed := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(seed)
	tickerJob := time.NewTicker(2 * time.Second) // Increase amount each 2 seconds

	amount := r.Intn(100)
	myJob := MainJob{
		ProcessInfo:       processInfo,
		NetworkInfo:       network,
		Data:              models.ProcessData{Amount: amount, Mu: sync.Mutex{}},
		updateStateChan:   updateStateChan,
		processMessageOut: processMessageOut,
		processMessageIn:  processMessageIn,
		ticker:            tickerJob,
		saveGlobalState:   saveGlobalState,
	}
	go myJob.MockJob()
	go myJob.ExecuteTaskList(taskList, quit)
	return &myJob
}
