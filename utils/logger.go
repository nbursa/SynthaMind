package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Ensure logs directory and file are initialized
func ensureLogFile() (*os.File, error) {
	// Ensure logs directory exists
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		return nil, fmt.Errorf("❌ Error creating logs directory: %v", err)
	}

	// Open log file
	file, err := os.OpenFile("logs/tasks.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("❌ Error opening log file: %v", err)
	}
	return file, nil
}

// LogTaskExecution logs completed task execution
func LogTaskExecution(task Task, duration time.Duration) {
	logFile, err := ensureLogFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close() // Always close file after logging

	log.SetOutput(logFile)
	logMsg := fmt.Sprintf("[%s] Task %d executed | Priority: %d | Duration: %v | Data: %s",
		time.Now().Format("2006-01-02 15:04:05"), task.ID, task.Priority, duration, task.Data)

	log.Println(logMsg)
	fmt.Println("📝 Task Logged:", logMsg)
}

// LogTaskExpiry logs removed expired tasks
func LogTaskExpiry(task Task) {
	logFile, err := ensureLogFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close() // Always close file after logging

	log.SetOutput(logFile)
	logMsg := fmt.Sprintf("[%s] Task %d expired & removed | Data: %s",
		time.Now().Format("2006-01-02 15:04:05"), task.ID, task.Data)

	log.Println(logMsg)
	fmt.Println("🗑️ Expired Task Logged:", logMsg)
}
