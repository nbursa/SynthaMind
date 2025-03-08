package main

import (
	"evolvai/taskmanager"
	"evolvai/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	// 📌 If in report mode, show valid tasks only
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

	// ⏳ Wait before exit to allow processing (tasks expire naturally)
	time.Sleep(5 * time.Second) // ✅ Ensures report checks still work
}
