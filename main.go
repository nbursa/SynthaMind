// package main

// import (
// 	"evolvai/taskmanager"
// 	"evolvai/utils"
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	// 📌 If in report mode, show valid tasks only
// 	if len(os.Args) > 1 && os.Args[1] == "report" {
// 		utils.ReportTaskStats()
// 		return
// 	}

// 	// 🚀 Start EvolvAI Task Manager
// 	fmt.Println("🚀 EvolvAI Task Manager Starting...")
// 	taskmanager.StartTaskManager()

// 	// 📌 Define test tasks to verify AI agent memory
// 	testTasks := []string{
// 		"System error detected",
// 		"Server error occurred",
// 		"Memory usage warning",
// 		"Memory leak detected",
// 		"Critical update required",
// 		"Urgent security patch needed",
// 		"Self-awareness data stored",
// 		"Pattern recognition triggered",
// 		"AI behavior anomaly detected",
// 	}

// 	// 🔄 Add predefined test tasks sequentially
// 	for _, taskText := range testTasks {
// 		fmt.Printf("➕ Adding Task: %s\n", taskText)
// 		taskmanager.AddTask(taskText)
// 		time.Sleep(1 * time.Second)
// 	}

// 	// ✅ **Fix: Properly wait before manual input mode**
// 	fmt.Println("\n🔵 **Manual Input Mode Activated**: Type a task, or type 'exit' to quit.")

// 	// **Ensure main waits properly for user input**
// 	var taskText string
// 	for {
// 		fmt.Print("➕ Enter a new task (or type 'exit' to quit): ")

// 		// Block for user input
// 		_, err := fmt.Scanln(&taskText)
// 		if err != nil {
// 			fmt.Println("⚠️ Input error. Try again.")
// 			continue
// 		}

// 		// **Exit condition**
// 		if taskText == "exit" {
// 			fmt.Println("🔴 Exiting manual task input mode...")
// 			break
// 		}

//			// **Add the user-provided task**
//			taskmanager.AddTask(taskText)
//		}
//	}
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

	// Start the task manager without waiting for user input
	taskmanager.StartTaskManager()

	// 📌 Define test tasks to verify AI agent memory
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

	// 🔄 Add predefined test tasks sequentially
	for _, taskText := range testTasks {
		fmt.Printf("➕ Adding Task: %s\n", taskText)
		taskmanager.AddTask(taskText)
		time.Sleep(1 * time.Second) // Simulate task arrival delay
	}

	// ⏳ Wait before exit to allow task processing to complete
	time.Sleep(10 * time.Second) // ✅ Ensures all tasks are processed before exit
}

