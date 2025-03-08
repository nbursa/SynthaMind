package taskmanager

import (
	"fmt"
	"sort"
	"time"

	"evolvai/agents"
	"evolvai/utils"
)

// Task queue (for AI task processing)
var taskQueue []utils.Task

// Task counter
var taskCounter int

// Expiry duration for tasks
const taskExpiry = 5 * time.Minute

// ✅ Initialize AI Agents correctly
var hippocampusAgent = agents.NewHippocampusAgent()
var cortexAgent = agents.NewCortexAgent(hippocampusAgent) // 🔥 Cortex now learns from memory!
var amygdalaAgent = agents.NewAmygdalaAgent()
var thalamusAgent = agents.NewThalamusAgent()
var executorAgent = agents.NewExecutorAgent()

// StartTaskManager initializes AI task processing
func StartTaskManager() {
	fmt.Println("🧠 AI Task Manager Running...")

	// ✅ Retrieve past tasks from HippocampusAgent (Long-Term Memory)
	pastTasks, err := hippocampusAgent.RetrieveMemory("all")
	if err == nil {
		thalamusAgent.LearnFromMemory(pastTasks) // ✅ Thalamus learns from past tasks
	}

	go processTasks()
}

// AddTask adds new tasks to the system, prioritizing HIGH-priority tasks first
func AddTask(data string) {
	taskCounter++
	newTask := utils.Task{
		ID:        taskCounter,
		Data:      data,
		Timestamp: time.Now(),
	}

	// ✅ Assign priority using Amygdala AI Agent
	amygdalaAgent.ProcessTask(&newTask)

	// ✅ Store task in HippocampusAgent
	hippocampusAgent.ProcessTask(&newTask)

	// Add task to queue
	taskQueue = append(taskQueue, newTask)

	// Sort queue by priority (high-priority tasks first)
	sort.Slice(taskQueue, func(i, j int) bool {
		return taskQueue[i].Priority > taskQueue[j].Priority
	})

	// Print updated queue for debugging
	printTaskQueue()
}

// processTasks continuously handles AI tasks (higher priority first)
func processTasks() {
	for {
		if len(taskQueue) > 0 {
			removeExpiredTasks() // ✅ Remove expired tasks before processing

			if len(taskQueue) == 0 {
				time.Sleep(1 * time.Second)
				continue
			}

			// Process next task
			task := taskQueue[0]
			taskQueue = taskQueue[1:] // Remove from queue

			startTime := time.Now()
			fmt.Printf("🟢 Processing Task %d (Priority: %d): %s\n", task.ID, task.Priority, task.Data)

			// ✅ Step 1: Thalamus filters the task first
			thalamusAgent.ProcessTask(&task)

			// ✅ Step 2: If task passes, Amygdala prioritizes it
			amygdalaAgent.ProcessTask(&task)

			// ✅ Step 3: Cortex processes task memory & reasoning
			cortexAgent.ProcessTask(&task) // 🔥 Now Cortex learns from past tasks!

			// ✅ Step 4: Executor performs final execution
			executorAgent.ProcessTask(task.Data)

			// Simulate processing delay
			taskDuration := time.Since(startTime)
			fmt.Printf("⏳ Task %d executed in %v\n", task.ID, taskDuration)
			fmt.Printf("📊 Remaining Tasks in Queue: %d\n", len(taskQueue))

			utils.LogTaskExecution(task, taskDuration)
			time.Sleep(2 * time.Second)
		} else {
			time.Sleep(1 * time.Second) // Wait before checking queue again
		}
	}
}

// ✅ Remove expired tasks from queue
func removeExpiredTasks() {
	now := time.Now()
	filteredQueue := []utils.Task{}

	for _, task := range taskQueue {
		if now.Sub(task.Timestamp) < taskExpiry {
			filteredQueue = append(filteredQueue, task) // Keep valid tasks
		} else {
			fmt.Printf("🗑️ Expired Task Removed: %d (%s)\n", task.ID, task.Data)
			utils.LogTaskExpiry(task) // ✅ Log expiry
		}
	}

	taskQueue = filteredQueue
}

// ✅ Debug function: Print current queue
func printTaskQueue() {
	fmt.Println("📌 Current Task Queue (Priority Order):")
	for _, task := range taskQueue {
		fmt.Printf("🔹 Task %d (Priority: %d): %s\n", task.ID, task.Priority, task.Data)
	}
}
