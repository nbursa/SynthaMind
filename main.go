// 🚀 SynthaMind - Task Manager Runner
// This script starts the SynthaMind Task Manager and processes AI-driven tasks.

package main

import (
	"evolvai/taskmanager"
	"evolvai/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	// 📌 Check if report mode is enabled via CLI argument
	if len(os.Args) > 1 && os.Args[1] == "report" {
		utils.ReportTaskStats()
		return
	}

	// 🚀 Start the SynthaMind Task Manager
	fmt.Println("🚀 SynthaMind Task Manager Starting...")
	taskmanager.StartTaskManager()

	// 📌 Define test tasks for AI task processing
	testTasks := []string{
		"System error detected",
		"Server error occurred",
		"Memory usage warning",
		"Memory leak detected",
		"Critical update required",
		"Urgent security patch needed",
		"Self-awareness data stored",
		"Pattern recognition triggered",
		"AI behavior anomaly detected",
	}

	// 🔄 Add predefined test tasks with simulated delays
	for _, taskText := range testTasks {
		fmt.Printf("➕ Adding Task: %s\n", taskText)
		taskmanager.AddTask(taskText)
		time.Sleep(1 * time.Second) // Simulating task arrival over time
	}

	// ⏳ Allow time for task processing before entering manual input mode
	fmt.Println("\n🔵 **Manual Input Mode Activated**: Type a task, or type 'exit' to quit.")

	// 📝 Manual task input loop
	for {
		fmt.Print("➕ Enter a new task (or type 'exit' to quit): ")

		// Read user input
		var taskText string
		_, err := fmt.Scanln(&taskText)
		if err != nil {
			fmt.Println("⚠️ Input error. Please try again.")
			continue
		}

		// 🚪 Exit condition
		if taskText == "exit" {
			fmt.Println("🔴 Exiting manual task input mode...")
			break
		}

		// ➕ Add user-defined task to task manager
		taskmanager.AddTask(taskText)
	}

	// ✅ Ensure a graceful shutdown
	fmt.Println("🔻 Shutting down SynthaMind Task Manager...")
	time.Sleep(2 * time.Second) // Simulate proper cleanup before exit
}
