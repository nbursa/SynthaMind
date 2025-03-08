package agents

import (
	"evolvai/utils"
	"fmt"
	"strings"
)

// AmygdalaAgent specializes in analyzing and prioritizing tasks.
type AmygdalaAgent struct {
	*BaseAgent
	memory map[string]utils.TaskPriority // ✅ Store past task data
}

// NewAmygdalaAgent initializes the AI agent for Amygdala
func NewAmygdalaAgent() *AmygdalaAgent {
	return &AmygdalaAgent{
		BaseAgent: NewAgent("Amygdala"),
		memory:    make(map[string]utils.TaskPriority), // ✅ Initialize memory storage
	}
}

// ProcessTask analyzes tasks and assigns priority.
func (a *AmygdalaAgent) ProcessTask(task *utils.Task) {
	fmt.Printf("🟠 AI Amygdala Agent processing: %s\n", task.Data)

	// ✅ Check if task was seen before
	if priority, exists := a.memory[task.Data]; exists {
		task.Priority = priority // ✅ Use learned priority
		fmt.Printf("🔁 AI Amygdala RECOGNIZED task, setting previous priority: %d\n", priority)
	} else {
		// ✅ If new task, analyze it
		a.analyzeTask(task)
		// ✅ Store decision in memory
		a.memory[task.Data] = task.Priority
	}
}

// ✅ Task Analysis Logic (adaptive)
func (a *AmygdalaAgent) analyzeTask(task *utils.Task) {
	// Convert task data to lowercase for case-insensitive comparison
	taskData := strings.ToLower(task.Data)

	// Assign priority based on keywords
	switch {
	case strings.Contains(taskData, "error"),
		strings.Contains(taskData, "critical"):
		task.Priority = utils.High
		fmt.Println("🔴 Task is critical! Assigning high priority.")
	case strings.Contains(taskData, "urgent"):
		task.Priority = utils.High // Assign high priority to urgent tasks
		fmt.Println("🔴 Task is urgent! Assigning high priority.")
	case strings.Contains(taskData, "update"),
		strings.Contains(taskData, "memory"):
		task.Priority = utils.Medium
		fmt.Println("🟠 Task is a warning or update. Assigning medium priority.")
	case strings.Contains(taskData, "self-awareness"),
		strings.Contains(taskData, "pattern"):
		task.Priority = utils.Low
		fmt.Println("🟢 Task is low priority (e.g., self-awareness).")
	default:
		// Default case for unknown tasks
		task.Priority = utils.Low
		fmt.Println("🟢 Task is not recognized. Assigning low priority.")
	}

	// Log the decision made by Amygdala
	fmt.Printf("🔹 AI Amygdala assigned priority %d to Task %d\n", task.Priority, task.ID)
}

