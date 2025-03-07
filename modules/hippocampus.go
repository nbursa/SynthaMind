package modules

import (
	"evolvai/utils"
	"fmt"
)

// Memory storage (temporary)
var memoryStorage = make(map[int]utils.Task)

// HippocampusStore saves knowledge in memory
func HippocampusStore(task utils.Task) {
	fmt.Printf("💾 Storing Task %d in memory: %s\n", task.ID, task.Data)

	// Debug: Print before storing
	fmt.Printf("🛠️ Current Memory Before: %+v\n", memoryStorage)

	// Store task
	memoryStorage[task.ID] = task

	// Debug: Print after storing
	fmt.Printf("✅ Task %d stored. Updated Memory: %+v\n", task.ID, memoryStorage)
}

// HippocampusRetrieve checks if the task is already known
func HippocampusRetrieve(task utils.Task) *utils.Task {
	// Debug: Print full memory storage before searching
	fmt.Printf("🔎 Hippocampus Searching for Task %d. Current Memory: %+v\n", task.ID, memoryStorage)

	if storedTask, exists := memoryStorage[task.ID]; exists {
		fmt.Printf("🔍 Hippocampus Found Memory for Task %d: %+v\n", task.ID, storedTask)
		return &storedTask
	}

	fmt.Println("❌ Hippocampus: No memory found for Task", task.ID)
	return nil
}
