/**
 * SynthaMind (c) 2025 Nenad Bursać
 * Licensed under AGPLv3 - https://www.gnu.org/licenses/agpl-3.0.html
 * Attribution required. Commercial use prohibited without permission.
 * See LICENSE file or https://nenadbursac.com/contact
 */

// SynthaMind - Task Manager Runner
// This script starts the SynthaMind Task Manager and processes AI-driven tasks.

package main

import (
	"fmt"
	"os"
	"synthamind/taskmanager"
	"synthamind/utils"
	"time"
)

func main() {
	//  Check if report mode is enabled via CLI argument
	if len(os.Args) > 1 && os.Args[1] == "report" {
		utils.ReportTaskStats()
		return
	}

	//  Start the SynthaMind Task Manager
	fmt.Println("🚀 SynthaMind Task Manager Starting...")
	taskmanager.StartTaskManager()

	//  Define test tasks for AI task processing
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

	// Add predefined test tasks with simulated delays
	for _, taskText := range testTasks {
		fmt.Printf("➕ Adding Task: %s\n", taskText)
		taskmanager.AddTask(taskText)
		time.Sleep(1 * time.Second) // Simulating task arrival over time
	}

	//  Ensure a graceful shutdown
	fmt.Println("🔻 Shutting down SynthaMind Task Manager...")
	time.Sleep(2 * time.Second) // Simulate proper cleanup before exit
}
