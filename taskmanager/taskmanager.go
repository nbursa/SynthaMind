package taskmanager

import (
	"fmt"
	"sort"
	"time"

	"evolvai/modules"
	"evolvai/utils"
)

// Task queue (priority-based)
var taskQueue []utils.Task

// Task counter
var taskCounter int

// StartTaskManager initializes AI task processing
func StartTaskManager() {
	fmt.Println("🧠 AI Task Manager Running...")
	go processTasks()
}

// AddTask adds new tasks, ensuring priority-based sorting
func AddTask(data string) {
	taskCounter++
	newTask := utils.Task{ID: taskCounter, Data: data}

	// Assign priority using Amygdala
	modules.AmygdalaAnalyze(&newTask)

	// Add task to queue
	taskQueue = append(taskQueue, newTask)

	// Sort queue by priority (higher-priority tasks first)
	sort.SliceStable(taskQueue, func(i, j int) bool {
		return taskQueue[i].Priority > taskQueue[j].Priority
	})

	fmt.Printf("📌 Added Task %d (Priority: %d): %s\n", newTask.ID, newTask.Priority, newTask.Data)
}

// processTasks continuously handles AI tasks (higher priority first)
// func processTasks() {
// 	for {
// 		if len(taskQueue) > 0 {
// 			// Extract task from queue
// 			task := taskQueue[0]
// 			taskQueue = taskQueue[1:] // Remove from queue

// 			fmt.Printf("🟢 Processing Task %d (Priority: %d): %s\n", task.ID, task.Priority, task.Data)

// 			// Execute task through AI modules
// 			go modules.ThalamusFilter(task)
// 			go modules.Executor(task)

// 			time.Sleep(2 * time.Second) // Simulate processing delay
// 		} else {
// 			time.Sleep(1 * time.Second) // Wait before checking queue again
// 		}
// 	}
// }
// func processTasks() {
// 	for {
// 		if len(taskQueue) > 0 {
// 			// Log queue state before processing
// 			fmt.Println("📌 Current Task Queue (Priority Order):")
// 			for _, task := range taskQueue {
// 				fmt.Printf("🔹 Task %d (Priority: %d): %s\n", task.ID, task.Priority, task.Data)
// 			}

// 			// Process the highest priority task first
// 			task := taskQueue[0]
// 			taskQueue = taskQueue[1:] // Remove first task from queue

// 			fmt.Printf("🟢 Processing Task %d (Priority: %d): %s\n", task.ID, task.Priority, task.Data)

// 			go modules.ThalamusFilter(task)
// 			go modules.Executor(task)

// 			time.Sleep(2 * time.Second) // Simulate processing delay
// 		} else {
// 			time.Sleep(1 * time.Second) // Wait before checking queue again
// 		}
// 	}
// }
func processTasks() {
	for {
		if len(taskQueue) > 0 {
			// Log queue state before processing
			fmt.Println("📌 Current Task Queue (Priority Order):")
			for _, task := range taskQueue {
				fmt.Printf("🔹 Task %d (Priority: %d): %s\n", task.ID, task.Priority, task.Data)
			}

			// Pick highest-priority task
			task := taskQueue[0]
			taskQueue = taskQueue[1:]

			// Track execution time
			startTime := time.Now()

			fmt.Printf("🟢 Processing Task %d (Priority: %d): %s\n", task.ID, task.Priority, task.Data)

			// Execute task
			go modules.ThalamusFilter(task)
			go modules.Executor(task)

			// Log execution time
			executionTime := time.Since(startTime)
			fmt.Printf("⏳ Task %d executed in %s\n", task.ID, executionTime)

			// Log remaining queue size
			fmt.Printf("📊 Remaining Tasks in Queue: %d\n", len(taskQueue))

			time.Sleep(2 * time.Second) // Simulate processing delay
		} else {
			time.Sleep(1 * time.Second) // Wait before checking queue again
		}
	}
}


