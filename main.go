package main

import (
	"evolvai/taskmanager"
	"evolvai/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	// 📌 Check if report mode is requested
	if len(os.Args) > 1 && os.Args[1] == "report" {
		utils.ReportTaskStats()
		return
	}

	// 🚀 Start EvolvAI Task Manager
	fmt.Println("🚀 EvolvAI Task Manager Starting...")
	taskmanager.StartTaskManager()

	// 📌 Define test tasks
	testTasks := []string{
		"System error detected", "Memory usage warning", "Critical update required",
		"Self-awareness data stored", "Pattern recognition triggered",
	}

	// 🔄 Add tasks sequentially with a delay
	for _, taskText := range testTasks {
		fmt.Printf("➕ Adding Task: %s\n", taskText) // ✅ Log added tasks
		taskmanager.AddTask(taskText)
		time.Sleep(1 * time.Second) // Simulate task arrival delay
	}

	// ⏳ Wait before exit to allow processing
	time.Sleep(10 * time.Second) // ✅ Ensure enough time for processing
}
