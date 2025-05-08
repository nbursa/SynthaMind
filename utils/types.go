// Package utils defines core types for SynthaMind tasks.
package utils

import "time"

// TaskPriority defines priority levels for tasks.
type TaskPriority int

const (
	High   TaskPriority = 3
	Medium TaskPriority = 2
	Low    TaskPriority = 1
)

// Task represents a unit of work in the AI task manager.
type Task struct {
	ID        int          `json:"id"`
	Data      string       `json:"data"`
	Vector    []float32    `json:"vector"`
	Priority  TaskPriority `json:"priority"`
	Timestamp time.Time    `json:"timestamp"` //  Used for task expiration tracking
}

// TaskVector is used for storing task embeddings in ChromaDB.
type TaskVector struct {
	ID       int       `json:"id"`
	TaskName string    `json:"task_name"`
	Vector   []float32 `json:"vector"`
}
