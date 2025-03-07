package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// EnsureChromaCollection creates the collection if it doesn't exist
func EnsureChromaCollection() error {
	url := "http://127.0.0.1:8000/api/v1/collections"

	payload, err := json.Marshal(map[string]interface{}{
		"name": "tasks",
		"metadata": map[string]interface{}{
			"description": "Collection for task vectors",
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
		return fmt.Errorf("failed to create ChromaDB collection: %s", resp.Status)
	}

	log.Println("✅ ChromaDB collection 'tasks' created successfully")
	return nil
}

// InitChroma ensures the collection exists before use
func InitChroma() error {
	if err := EnsureChromaCollection(); err != nil {
		return err
	}
	return nil
}