package modules

import (
	"fmt"
	"synthamind/utils"
	"time"
)

// Define memory limits
const maxShortTermMemory = 10
const taskExpiryTime = 30 * time.Second

// TaskEntry represents a memory entry with a timestamp
type TaskEntry struct {
	Task      utils.Task
	Timestamp time.Time
}

// Memory storage
var memoryStorage = make(map[int]TaskEntry)

// ✅ Stores only valid tasks, automatically removes expired ones
func HippocampusStore(task utils.Task) {
	// ✅ Remove expired tasks first (not logged, just naturally gone)
	cleanupShortTermMemory()

	// ✅ Add new valid task
	entry := TaskEntry{
		Task:      task,
		Timestamp: time.Now(),
	}
	memoryStorage[task.ID] = entry
	fmt.Printf("💾 Task %d stored: %s\n", task.ID, task.Data)
}

// ✅ Retrieves only **non-expired** tasks
func HippocampusRetrieve(task utils.Task) *utils.Task {
	if entry, exists := memoryStorage[task.ID]; exists {
		// ✅ If expired, delete it and return nil
		if time.Since(entry.Timestamp) > taskExpiryTime {
			delete(memoryStorage, task.ID) // 🔥 Task disappears automatically
			return nil
		}
		return &entry.Task
	}
	return nil
}

// ✅ Removes expired tasks *without logging anything*
func cleanupShortTermMemory() {
	now := time.Now()
	for id, entry := range memoryStorage {
		if now.Sub(entry.Timestamp) > taskExpiryTime {
			delete(memoryStorage, id) // 🔥 Expired tasks vanish silently
		}
	}
}
