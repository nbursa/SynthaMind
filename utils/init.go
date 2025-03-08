// Package utils provides initialization functions for SynthaMind.
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// EnsureChromaCollection ensures the required collection exists in ChromaDB.
func EnsureChromaCollection() error {
	url := "http://127.0.0.1:8000/api/v1/collections"

	payload, err := json.Marshal(map[string]interface{}{
		"name": "tasks",
		"metadata": map[string]interface{}{
			"description": "Collection for task vectors",
		},
	})
	if err != nil {
		return fmt.Errorf("❌ Failed to encode JSON payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("❌ Failed to connect to ChromaDB: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		log.Println("⚠️ ChromaDB collection 'tasks' already exists. Skipping creation.")
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("❌ Failed to create ChromaDB collection: %s", resp.Status)
	}

	log.Println("✅ ChromaDB collection 'tasks' created successfully.")
	return nil
}

// InitChroma initializes ChromaDB by ensuring the required collection exists.
func InitChroma() error {
	return EnsureChromaCollection()
}
