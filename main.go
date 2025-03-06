package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"evolvai/taskmanager"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Pinecone API Key and Endpoint from environment
	apiKey := os.Getenv("PINECONE_API_KEY")
	if apiKey == "" {
		log.Fatal("PINECONE_API_KEY not set in .env file")
	}

	endpoint := os.Getenv("PINECONE_API_ENDPOINT")
	if endpoint == "" {
		log.Fatal("PINECONE_API_ENDPOINT not set in .env file")
	}

	// Pinecone namespace
	namespace := "evolvai-memory" // Adjust namespace as needed

	// Sample vector data (replace this with the task vector)
	vector := map[string]interface{}{
		"id": "1",
		"values": []float32{0.1, 0.2, 0.3, 0.4, 0.5},  // Sample vector data
		"metadata": map[string]interface{}{
			"text": "Sample task data",
		},
	}

	// Prepare the request payload
	payload := map[string]interface{}{
		"vectors": []interface{}{vector},
		"namespace": namespace,
	}

	// Convert payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshaling payload:", err)
	}

	// Send POST request to Pinecone API
	req, err := http.NewRequest("POST", endpoint+"/vectors/upsert", bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
	}

	// Set the authorization header with the API key
	req.Header.Set("Authorization", "Api-Key "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error with Pinecone API request: %s", resp.Status)
	} else {
		fmt.Println("Data successfully upserted to Pinecone")
	}

	// Start task manager
	fmt.Println("🚀 EvolvAI Task Manager Starting...")
	go taskmanager.StartTaskManager()

	// Simulate incoming knowledge tasks
	for {
		taskmanager.AddTask("New knowledge detected")
		time.Sleep(5 * time.Second)
	}
}
