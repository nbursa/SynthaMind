package main

import (
	"evolvai/taskmanager"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Start EvolvAI Task Manager
	fmt.Println("🚀 EvolvAI Task Manager Starting...")
	taskmanager.StartTaskManager() // Start task processing system

	// 🔹 Test tasks with different priority levels
	testTasks := []string{
		"Critical update required",  // High Priority
		"System error detected",     // High Priority
		"Memory usage warning",      // Medium Priority
		"Self-awareness data stored", // Low Priority
		"Pattern recognition triggered", // Low Priority
		"Disk space low warning",      // Medium Priority
		"High CPU temperature detected", // Medium Priority
		"Network latency spike detected", // Low Priority
		"Filesystem corruption warning", // Medium Priority
		"GPU overheating alert",        // High Priority
		"Unexpected power failure detected", // High Priority
	}

	// Add tasks to Task Manager
	for _, taskText := range testTasks {
		taskmanager.AddTask(taskText) // Tasks are automatically prioritized
		time.Sleep(500 * time.Millisecond) // Simulate task arrival rate
	}

	// Keep running to allow tasks to be processed
	time.Sleep(10 * time.Second)
}
