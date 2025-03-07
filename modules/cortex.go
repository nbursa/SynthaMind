package modules

import (
	"evolvai/chroma"
	"evolvai/utils"

	// "evolvai/utils/types"
	"fmt"
	"net/http"
	"time"
)

// CortexProcess handles AI learning and memory
func CortexProcess(task utils.Task) {
	fmt.Printf("🧠 Cortex Processing Task %d: %s\n", task.ID, task.Data)

	// Convert task text to a vector
	vector := GenerateVector(task.Data)

	// Ensure ChromaDB is ready before proceeding
	if err := chroma.EnsureChromaCollection(); err != nil {
		fmt.Println("❌ Skipping task processing due to ChromaDB unavailability.")
		return
	}

	// Step 1️⃣: Retrieve similar knowledge from ChromaDB
	similarTasks, err := chroma.SearchTaskInChroma(vector, 3) // Find top 3 similar tasks
	if err != nil {
		fmt.Println("❌ Retrieval failed, skipping this step.")
		return
	}

	// Step 2️⃣: If similar knowledge exists, print it
	if len(similarTasks) > 0 {
		fmt.Println("🔍 Similar knowledge found:")
		for _, similar := range similarTasks {
			fmt.Printf("📝 Similar Task: %s\n", similar.TaskName)
		}
	} else {
		fmt.Println("❌ No similar knowledge found, storing new entry...")
	}

	// Step 3️⃣: Store the new knowledge in ChromaDB
	chroma.AddTaskToChroma(utils.TaskVector{
		ID:       task.ID,
		TaskName: task.Data,
		Vector:   vector, // Use generated vector
	})
	fmt.Println("✅ Task stored in memory (ChromaDB).")
}

// GenerateVector converts text to a vector (placeholder)
func GenerateVector(data string) []float32 {
	// Placeholder for real AI embedding model (later use OpenAI, BERT, etc.)
	return []float32{0.1, 0.2, 0.3} // Temporary fixed vector
}

// WaitForChromaDB actively checks if ChromaDB is ready
func WaitForChromaDB() bool {
	fmt.Println("⏳ Waiting for ChromaDB to be available...")

	for i := 0; i < 5; i++ { // Try 5 times
		resp, err := http.Get("http://127.0.0.1:8000/api/v1/heartbeat")
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("✅ ChromaDB is ready!")
			return true
		}
		fmt.Println("⚠️ ChromaDB not available, retrying...")
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	fmt.Println("❌ ChromaDB is still unavailable after retries.")
	return false
}