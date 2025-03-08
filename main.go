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

	// Start Task Manager
	taskmanager.StartTaskManager()

	// Add tasks with mixed priorities
	taskmanager.AddTask("System error detected")      // High Priority
	taskmanager.AddTask("Memory usage warning")       // Medium Priority
	taskmanager.AddTask("Critical update required")   // High Priority
	taskmanager.AddTask("Self-awareness data stored") // Low Priority
	taskmanager.AddTask("Pattern recognition triggered") // Low Priority

	// Allow tasks to be processed
	time.Sleep(15 * time.Second)
}
