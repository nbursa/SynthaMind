package modules

import (
	"evolvai/chroma"
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

// Cleanup short-term memory by moving expired tasks to Neocortex (ChromaDB)
func cleanupShortTermMemory() {
	if len(memoryStorage) > maxShortTermMemory {
		fmt.Println("🧹 Hippocampus Cleanup: Moving old memories to Neocortex...")

		// Find the oldest task
		var oldestID int
		var oldestTime time.Time

		for id, entry := range memoryStorage {
			if oldestTime.IsZero() || entry.Timestamp.Before(oldestTime) {
				oldestID = id
				oldestTime = entry.Timestamp
			}
		}

		// Move oldest task to Neocortex before deleting
		if oldestID != 0 {
			oldestTask := memoryStorage[oldestID].Task
			NeocortexStore(oldestTask) // Store in Long-Term Memory
			delete(memoryStorage, oldestID)
			fmt.Printf("🗑️ Moved Task %d to Neocortex and removed from short-term memory.\n", oldestID)
		}
	}
}

// NeocortexStore moves long-term knowledge into ChromaDB
func NeocortexStore(task utils.Task) {
	fmt.Printf("💾 Moving Task %d to Neocortex (Long-Term Memory): %s\n", task.ID, task.Data)

	// Ensure ChromaDB collection exists
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Failed to connect to ChromaDB for Long-Term Memory.")
		return
	}

	// Convert task to vector
	vector := GenerateVector(task.Data)

	// Store in ChromaDB
	err = chroma.AddTaskToChroma(collectionID, utils.TaskVector{
		ID:       task.ID,
		TaskName: task.Data,
		Vector:   vector,
	})
	if err != nil {
		fmt.Println("❌ Failed to store task in Neocortex (ChromaDB).")
		return
	}

	fmt.Println("✅ Task successfully stored in Neocortex (Long-Term Memory).")
}
