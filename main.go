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

	// Simulate incoming tasks
	for i := 1; i <= 3; i++ {
		task := utils.Task{
			ID:   i,
			Data: fmt.Sprintf("New knowledge entry %d", i),
		}
		modules.ThalamusFilter(task) // Send task through the pipeline
		time.Sleep(3 * time.Second) // Simulate time between tasks
	}
}