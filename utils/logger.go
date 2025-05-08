// Package utils provides logging functionality for SynthaMind.
package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logFile *os.File //  Declared globally for proper file handling.

func init() {
	// Ensure the "logs" directory exists before writing logs.
	os.MkdirAll("logs", os.ModePerm)

	var err error
	logFile, err = os.OpenFile("logs/tasks.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Error initializing log file:", err)
		return
	}
	log.SetOutput(logFile)
}

// LogTaskExecution logs the execution of a completed task.
func LogTaskExecution(task Task, duration time.Duration) {
	logMsg := fmt.Sprintf("[%s] Task %d executed | Priority: %d | Duration: %v | Data: %s",
		time.Now().Format("2006-01-02 15:04:05"), task.ID, task.Priority, duration, task.Data)
	log.Println(logMsg)
	fmt.Println("📝 Task Logged:", logMsg)
}

// LogTaskExpiry logs tasks that were removed due to expiration.
func LogTaskExpiry(task Task) {
	logMsg := fmt.Sprintf("[%s] Task %d expired & removed | Data: %s",
		time.Now().Format("2006-01-02 15:04:05"), task.ID, task.Data)

	log.Println(logMsg)
	fmt.Println("🗑️ Expired Task Logged:", logMsg)
}
