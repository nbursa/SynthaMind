package modules

import (
	"crypto/sha256"
	"encoding/binary"
	"evolvai/chroma"
	"evolvai/utils"
	"fmt"
	"net/http"
	"time"
)

// CortexProcess processes tasks and integrates memory storage
func CortexProcess(task utils.Task) {
	fmt.Printf("🧠 Cortex Processing Task %d: %s\n", task.ID, task.Data)

	// 🔍 Step 1: Check if task exists in Hippocampus (memory)
	memory := HippocampusRetrieve(task)
	if memory != nil {
		fmt.Println("🔁 Using past memory instead of reprocessing.")
		return
	}

	// 🔹 Step 2: Convert task text to a vector
	vector := GenerateVector(task.Data)

	// 🔹 Step 3: Ensure ChromaDB collection exists
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Skipping task processing due to ChromaDB unavailability.")
		return
	}

	// 🔹 Step 4: Store new knowledge
	err = chroma.AddTaskToChroma(collectionID, utils.TaskVector{
		ID:       task.ID,
		TaskName: task.Data,
		Vector:   vector,
	})
	if err != nil {
		fmt.Println("❌ Failed to store task in ChromaDB.")
		return
	}

	// 🔹 Step 5: Store in Hippocampus memory
	HippocampusStore(task)
	fmt.Println("✅ New knowledge stored in Hippocampus.")
}

// GenerateVector converts text into a hashed float32 vector
func GenerateVector(data string) []float32 {
	hash := sha256.Sum256([]byte(data))
	vector := make([]float32, 3)

	for i := 0; i < 3; i++ {
		vector[i] = float32(binary.BigEndian.Uint32(hash[i*8 : (i+1)*8])) / 1e9
	}
	return vector
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
		time.Sleep(2 * time.Second)
	}

	fmt.Println("❌ ChromaDB is still unavailable after retries.")
	return false
}
