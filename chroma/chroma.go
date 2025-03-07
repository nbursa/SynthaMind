package chroma

import (
	"bytes"
	"encoding/json"
	"evolvai/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func EnsureChromaCollection() error {
	url := "http://127.0.0.1:8000/api/v1/collections"

	// Fetch list of collections
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Failed to fetch collections:", err)
		return fmt.Errorf("failed to fetch collections: %w", err)
	}
	defer resp.Body.Close()

	// ChromaDB returns an **array**, not a map
	var existingCollections []map[string]interface{} // ✅ Change to slice of maps
	if err := json.NewDecoder(resp.Body).Decode(&existingCollections); err != nil {
		fmt.Println("❌ Failed to decode ChromaDB response:", err)
		return fmt.Errorf("failed to decode collection response: %w", err)
	}

	// Debugging: Print API response
	fmt.Println("🔍 ChromaDB Collection Response:", existingCollections)

	// Check if "tasks" collection is present
	for _, col := range existingCollections {
		if col["name"] == "tasks" {
			fmt.Println("✅ ChromaDB collection 'tasks' exists.")
			return nil // ✅ Collection exists, return success
		}
	}

	// If collection is missing, create it
	fmt.Println("🔵 Creating new ChromaDB collection: 'tasks'...")
	createURL := "http://127.0.0.1:8000/api/v1/collections"
	payload, _ := json.Marshal(map[string]interface{}{
		"name": "tasks",
	})

	resp, err = http.Post(createURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ChromaDB collection creation error: %s", resp.Status)
	}

	fmt.Println("✅ ChromaDB collection 'tasks' created successfully.")
	return nil
}

func AddTaskToChroma(task utils.TaskVector) error {
	// Ensure collection exists before inserting data
	if err := EnsureChromaCollection(); err != nil {
		return err
	}

	// Get ChromaDB collection ID instead of name
	collectionID := "3f02fd34-386f-42d2-b4cb-de8215959b04" // Hardcoded for now, should be retrieved dynamically

	url := "http://127.0.0.1:8000/api/v1/collections/" + collectionID + "/upsert"

	payload, err := json.Marshal(map[string]interface{}{
		"ids":        []string{strconv.Itoa(task.ID)},
		"embeddings": [][]float32{task.Vector},
		"metadatas": []map[string]interface{}{
			{"task_name": task.TaskName},
		},
	})
	if err != nil {
		return err
	}

	// Debugging: Print the JSON payload before sending
	fmt.Println("📤 Sending Payload to ChromaDB:", string(payload))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body for debugging
	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)
	fmt.Println("📥 ChromaDB Response:", responseBody)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ChromaDB error: %s | Response: %v", resp.Status, responseBody)
	}

	log.Printf("✅ Task '%s' stored in ChromaDB", task.TaskName)
	return nil
}

func SearchTaskInChroma(queryVector []float32, topK int) ([]utils.TaskVector, error) {
	// Ensure collection exists before searching
	if err := EnsureChromaCollection(); err != nil {
		return nil, err
	}

	// Use the correct collection ID (retrieved dynamically or hardcoded for now)
	collectionID := "3f02fd34-386f-42d2-b4cb-de8215959b04" // 🔹 Replace "tasks" with the actual collection ID

	url := "http://127.0.0.1:8000/api/v1/collections/" + collectionID + "/query"

	payload, err := json.Marshal(map[string]interface{}{
		"query_embeddings": [][]float32{queryVector},
		"n_results":        topK,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// **Ensure required fields exist**
	idsRaw, idsOk := result["ids"].([]interface{})
	metadatasRaw, metaOk := result["metadatas"].([]interface{})

	if !idsOk || !metaOk || len(idsRaw) == 0 || len(metadatasRaw) == 0 {
		return nil, fmt.Errorf("ChromaDB response missing 'ids' or 'metadatas' fields: %v", result)
	}

	// **Extract nested lists**
	idList, idListOk := idsRaw[0].([]interface{})
	metadataList, metadataListOk := metadatasRaw[0].([]interface{})

	if !idListOk || !metadataListOk {
		return nil, fmt.Errorf("unexpected response format from ChromaDB")
	}

	tasks := []utils.TaskVector{}
	for i := range idList {
		metadataMap, ok := metadataList[i].(map[string]interface{})
		if !ok {
			log.Printf("⚠️ Warning: Unexpected metadata format: %v", metadataList[i])
			continue
		}

		taskName, ok := metadataMap["task_name"].(string)
		if !ok {
			taskName = "Unknown"
		}

		tasks = append(tasks, utils.TaskVector{
			ID:       i + 1, // Assigning a numerical ID
			TaskName: taskName,
			Vector:   queryVector,
		})
	}

	return tasks, nil
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
		fmt.Printf("⚠️ ChromaDB not available (attempt %d). Error: %v\n", i+1, err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	fmt.Println("❌ ChromaDB is still unavailable after retries.")
	return false
}
