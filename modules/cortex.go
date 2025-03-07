package modules

import (
	"evolvai/utils"
	"fmt"
)

// CortexProcess handles AI learning and memory
func CortexProcess(task utils.Task) {
	fmt.Printf("🧠 Cortex Processing Task %d: %s\n", task.ID, task.Data)

	// Check if task should be stored in ChromaDB
	if ShouldStoreInMemory(task) {
		utils.AddTaskToChroma(utils.TaskVector{
			ID:       task.ID,
			TaskName: task.Data,
			Vector:   GenerateVector(task.Data), // Convert to vector
		})
		fmt.Println("✅ Task stored in memory (ChromaDB).")
	} else {
		fmt.Println("❌ Task not important enough to store.")
	}
}

// ShouldStoreInMemory determines if knowledge should be stored
func ShouldStoreInMemory(task utils.Task) bool {
	return len(task.Data) > 10 // Basic rule (later, use AI models)
}

// GenerateVector converts text to a vector (placeholder)
func GenerateVector(data string) []float32 {
	// Placeholder for real AI embedding model (later use OpenAI, BERT, etc.)
	return []float32{0.1, 0.2, 0.3} // Temporary fixed vector
}