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

// EnsureChromaCollection checks if the collection exists
func EnsureChromaCollection() error {
	// Add implementation for collection creation if needed
	return nil
}

// AddTaskToChroma stores a task vector in ChromaDB
func AddTaskToChroma(task utils.TaskVector) error {
	// Ensure collection exists before inserting data
	if err := EnsureChromaCollection(); err != nil {
		return err
	}

	url := "http://127.0.0.1:8000/api/v1/collections/tasks/upsert"

	payload, err := json.Marshal(map[string]interface{}{
		"ids":       []string{strconv.Itoa(task.ID)},
		"metadatas": []map[string]string{{"task_name": task.TaskName}},
		"embeddings": [][]float32{
			task.Vector,
		},
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ChromaDB error: %s", resp.Status)
	}

	log.Printf("✅ Task '%s' stored in ChromaDB", task.TaskName)
	return nil
}

// SearchTaskInChroma retrieves similar tasks based on the query vector
func SearchTaskInChroma(queryVector []float32, topK int) ([]utils.TaskVector, error) {
	url := "http://127.0.0.1:8000/api/v1/collections/tasks/query"

	payload, err := json.Marshal(map[string]interface{}{
		"collection_name": "tasks", // Fix: Use "collection_name" instead of "collection"
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

	// Ensure response contains required fields
	idsRaw, idsOk := result["ids"].([]interface{})
	metadatasRaw, metaOk := result["metadatas"].([]interface{})

	if !idsOk || !metaOk {
		return nil, fmt.Errorf("ChromaDB response missing 'ids' or 'metadatas' fields: %v", result)
	}

	tasks := []utils.TaskVector{}
	for i, id := range idsRaw {
		idFloat, ok := id.(float64)
		if !ok {
			log.Printf("⚠️ Warning: Unexpected ID format: %v", id)
			continue
		}

		metadataMap, ok := metadatasRaw[i].(map[string]interface{})
		if !ok {
			log.Printf("⚠️ Warning: Unexpected metadata format: %v", metadatasRaw[i])
			continue
		}

		taskName, ok := metadataMap["task_name"].(string)
		if !ok {
			log.Printf("⚠️ Warning: Missing 'task_name' in metadata: %v", metadataMap)
			taskName = "Unknown"
		}

		tasks = append(tasks, utils.TaskVector{
			ID:       int(idFloat),
			TaskName: taskName,
			Vector:   queryVector, 
		})
	}

	return tasks, nil
}