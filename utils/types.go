package utils

import "time"

// TaskPriority defines priority levels
type TaskPriority int

const (
	High   TaskPriority = 3
	Medium TaskPriority = 2
	Low    TaskPriority = 1
)

// Task structure used across the system
type Task struct {
	ID       int
	Data     string
	Vector   []float32
	Priority TaskPriority
    Timestamp time.Time // ✅ New field for expiry tracking
}

// TaskVector used for ChromaDB storage
type TaskVector struct {
	ID       int
	TaskName string
	Vector   []float32
}
