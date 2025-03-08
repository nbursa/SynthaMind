package modules

import (
	"evolvai/utils"
	"fmt"
	"strings"
)

// Executor handles task execution based on the task type.
func Executor(task utils.Task) {
	fmt.Printf("⚡ Executing Task %d: %s\n", task.ID, task.Data)

	// Define execution logic based on task content
	switch {
	case strings.Contains(strings.ToLower(task.Data), "error"):
		handleError(task)
	case strings.Contains(strings.ToLower(task.Data), "memory"):
		handleMemoryCheck(task)
	case strings.Contains(strings.ToLower(task.Data), "update"):
		handleSystemUpdate(task)
	case strings.Contains(strings.ToLower(task.Data), "self-awareness"):
		handleSelfAwareness(task)
	default:
		fmt.Println("⚪ No execution required for this task.")
	}
}

// handleError - Example: Log and escalate critical errors
func handleError(task utils.Task) {
	fmt.Println("🚨 Critical Error Detected! Logging and escalating:", task.Data)
	// Add future escalation logic here (send alert, etc.)
}

// handleMemoryCheck - Example: Log memory usage
func handleMemoryCheck(task utils.Task) {
	fmt.Println("💾 Checking system memory status...")
	// Add future memory management logic here
}

// handleSystemUpdate - Example: Process system update tasks
func handleSystemUpdate(task utils.Task) {
	fmt.Println("🔄 Processing system update task:", task.Data)
	// Future: Fetch updates, log changes, notify user
}

// handleSelfAwareness - Example: Log self-awareness updates
func handleSelfAwareness(task utils.Task) {
	fmt.Println("🧠 Processing self-awareness update:", task.Data)
	// Future: Compare old vs. new self-awareness data
}
