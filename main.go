package main

import (
	"evolvai/modules"
	"evolvai/utils"
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

	// Start the autonomous CortexBase loop
	go modules.CortexBase()

	// 🔹 Send test tasks that actually contain important keywords
	testTasks := []string{
		"System error detected",       // ✅ Matches "error"
		"Memory usage warning",        // ✅ Matches "memory" & "warning"
		"Critical update required",    // ✅ Matches "critical" & "update"
		"Self-awareness data stored",  // ✅ Matches "self-awareness"
		"Pattern recognition triggered", // ✅ Matches "pattern"
	}

	// Send tasks through Thalamus
	for i, taskText := range testTasks {
		task := utils.Task{
			ID:   i + 1,
			Data: taskText,
		}
		modules.ThalamusFilter(task) // Send task through the pipeline
		time.Sleep(1 * time.Second) // Simulate time between tasks
	}
}