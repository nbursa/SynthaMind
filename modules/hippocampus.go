package modules

import (
	"evolvai/utils"
	"fmt"
	"time"
)

// Define max tasks in short-term memory
const maxShortTermMemory = 10

// TaskEntry represents a memory entry with a timestamp
type TaskEntry struct {
	Task      utils.Task
	Timestamp time.Time
}

// Memory storage (temporary short-term memory)
var memoryStorage = make(map[int]TaskEntry)

// HippocampusStore saves knowledge in short-term memory
func HippocampusStore(task utils.Task) {
	entry := TaskEntry{
		Task:      task,
		Timestamp: time.Now(),
	}

	fmt.Printf("💾 Storing Task %d in short-term memory: %s\n", task.ID, task.Data)

	// Store task
	memoryStorage[task.ID] = entry

	// Perform cleanup if memory exceeds limit
	cleanupShortTermMemory()
}

// HippocampusRetrieve checks if a task is already known
func HippocampusRetrieve(task utils.Task) *utils.Task {
	fmt.Printf("🔎 Hippocampus Searching for Task %d...\n", task.ID)

	if entry, exists := memoryStorage[task.ID]; exists {
		fmt.Printf("🔍 Hippocampus Found Memory for Task %d: %+v\n", task.ID, entry.Task)
		return &entry.Task
	}

	fmt.Println("❌ Hippocampus: No memory found for Task", task.ID)
	return nil
}

// Cleanup short-term memory by keeping only the last N tasks
func cleanupShortTermMemory() {
	if len(memoryStorage) > maxShortTermMemory {
		fmt.Println("🧹 Hippocampus Cleanup: Removing oldest tasks...")

		// Find the oldest task
		var oldestID int
		var oldestTime time.Time

		for id, entry := range memoryStorage {
			if oldestTime.IsZero() || entry.Timestamp.Before(oldestTime) {
				oldestID = id
				oldestTime = entry.Timestamp
			}
		}

		// Remove oldest task
		if oldestID != 0 {
			delete(memoryStorage, oldestID)
			fmt.Printf("🗑️ Removed Task %d from short-term memory (expired)\n", oldestID)
		}
	}
}
