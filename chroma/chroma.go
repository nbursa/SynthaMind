package chroma

import (
	"bytes"
	"encoding/json"
	"evolvai/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// EnsureChromaCollection checks if the ChromaDB collection exists and returns its ID.
func EnsureChromaCollection() (string, error) {
	url := "http://127.0.0.1:8000/api/v1/collections"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Failed to fetch collections:", err)
		return "", fmt.Errorf("failed to fetch collections: %w", err)
	}
	defer resp.Body.Close()

	var existingCollections []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&existingCollections); err != nil {
		fmt.Println("❌ Failed to decode ChromaDB response:", err)
		return "", fmt.Errorf("failed to decode collection response: %w", err)
	}

	// Debugging: Print API response
	fmt.Println("🔍 ChromaDB Collection Response:", existingCollections)

	// Check if "tasks" collection exists
	for _, col := range existingCollections {
		if col["name"] == "tasks" {
			collectionID, ok := col["id"].(string)
			if !ok {
				return "", fmt.Errorf("collection found but has no valid ID")
			}
			fmt.Println("✅ ChromaDB collection 'tasks' exists with ID:", collectionID)
			return collectionID, nil
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
		return "", fmt.Errorf("failed to create collection: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ChromaDB collection creation error: %s", resp.Status)
	}

	// Decode response to get the newly created collection ID
	var createdCollection map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&createdCollection); err != nil {
		return "", fmt.Errorf("failed to decode collection creation response: %w", err)
	}

	collectionID, ok := createdCollection["id"].(string)
	if !ok {
		return "", fmt.Errorf("created collection but could not retrieve ID")
	}

	fmt.Println("✅ ChromaDB collection 'tasks' created successfully with ID:", collectionID)
	return collectionID, nil
}

// AddTaskToChroma stores a task in ChromaDB
func AddTaskToChroma(collectionID string, task utils.TaskVector) error {
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

// SearchTaskInChroma finds similar tasks using vector search in ChromaDB
func SearchTaskInChroma(collectionID string, queryVector []float32, topK int) ([]utils.TaskVector, error) {
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

	idsRaw, idsOk := result["ids"].([]interface{})
	metadatasRaw, metaOk := result["metadatas"].([]interface{})

	if !idsOk || !metaOk || len(idsRaw) == 0 || len(metadatasRaw) == 0 {
		return nil, fmt.Errorf("ChromaDB response missing 'ids' or 'metadatas' fields: %v", result)
	}

	tasks := []utils.TaskVector{}
	for i := range idsRaw {
		metadataArray, ok := metadatasRaw[i].([]interface{})
		if !ok || len(metadataArray) == 0 {
			fmt.Println("⚠️ Skipping task due to unexpected metadata structure:", metadatasRaw[i])
			continue
		}

		metadata, ok := metadataArray[0].(map[string]interface{})
		if !ok {
			fmt.Println("⚠️ Skipping task due to invalid metadata format:", metadataArray)
			continue
		}

		taskName, nameOk := metadata["task_name"].(string)
		if !nameOk {
			taskName = "Unknown"
		}

		tasks = append(tasks, utils.TaskVector{
			ID:       i + 1,
			TaskName: taskName,
			Vector:   queryVector,
		})
	}

	return tasks, nil
}
