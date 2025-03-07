package main

import (
	"evolvai/cortex"
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
	go cortex.CortexBase()

	// 🔹 Send test tasks that actually contain important keywords
	testTasks := []string{
		"System error detected",       // ✅ Matches "error"
		"Memory usage warning",        // ✅ Matches "memory" & "warning"
		"Critical update required",    // ✅ Matches "critical" & "update"
		"Self-awareness data stored",  // ✅ Matches "self-awareness"
		"Pattern recognition triggered", // ✅ Matches "pattern"
	}

	// Simulate incoming tasks
	// for i := 1; i <= 3; i++ {
	// 	task := utils.Task{
	// 		ID:   i,
	// 		Data: fmt.Sprintf("New knowledge entry %d", i),
	// 	}
	// 	modules.ThalamusFilter(task) // Send task through the pipeline
	// 	time.Sleep(3 * time.Second) // Simulate time between tasks
	// }
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