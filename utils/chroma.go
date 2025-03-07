package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// AddTaskToChroma stores a task vector in ChromaDB
func AddTaskToChroma(task TaskVector) error {
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