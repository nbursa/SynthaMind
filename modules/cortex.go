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

func CortexProcess(task utils.Task) {
	fmt.Printf("🧠 Cortex Processing Task %d: %s\n", task.ID, task.Data)

	// Convert task text to a vector
	vector := GenerateVector(task.Data)

	// Ensure ChromaDB collection exists and retrieve its ID
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Skipping task processing due to ChromaDB unavailability.")
		fmt.Println("📌 Debug Info: EnsureChromaCollection() Error →", err)
		return
	}

	// Step 1️⃣: Retrieve similar knowledge from ChromaDB
	similarTasks, err := chroma.SearchTaskInChroma(collectionID, vector, 3) // Pass collectionID
	if err != nil {
		fmt.Println("❌ Retrieval failed, skipping this step.")
		fmt.Println("📌 Debug Info: SearchTaskInChroma() Error →", err)
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
	err = chroma.AddTaskToChroma(collectionID, utils.TaskVector{ // Pass collectionID correctly
		ID:       task.ID,
		TaskName: task.Data,
		Vector:   vector,
	})
	if err != nil {
		fmt.Println("❌ Failed to store task in ChromaDB.")
		fmt.Println("📌 Debug Info: AddTaskToChroma() Error →", err)
		return
	}
	fmt.Println("✅ Task stored in memory (ChromaDB).")
}

// GenerateVector converts text into a hashed float32 vector
func GenerateVector(data string) []float32 {
	hash := sha256.Sum256([]byte(data))
	vector := make([]float32, 3) // Match the 3D placeholder size

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
