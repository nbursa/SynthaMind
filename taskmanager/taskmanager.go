package taskmanager

import (
	"fmt"
	"time"

	"evolvai/modules"
	"evolvai/utils"
)

// Task queue (for AI task processing)
var taskQueue = make(chan utils.Task, 10)

// Task counter
var taskCounter int

// StartTaskManager initializes AI task processing
func StartTaskManager() {
	fmt.Println("🧠 AI Task Manager Running...")
	go processTasks()
}

// AddTask sends new tasks to the AI system
func AddTask(data string) {
	taskCounter++
	taskQueue <- utils.Task{ID: taskCounter, Data: data}
}

// processTasks continuously handles AI tasks
func processTasks() {
	for task := range taskQueue {
		fmt.Printf("🟢 Processing Task %d: %s\n", task.ID, task.Data)

		go modules.ThalamusFilter(task)
		go modules.AmygdalaAnalyze(task)

		time.Sleep(2 * time.Second)
	}
}
