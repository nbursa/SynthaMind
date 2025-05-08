package agents

import (
	"fmt"
	"synthamind/chroma"
	"synthamind/utils"
)

// HippocampusAgent manages memory-related operations.
type HippocampusAgent struct {
	*BaseAgent
}

// NewHippocampusAgent initializes the Hippocampus AI agent.
func NewHippocampusAgent() *HippocampusAgent {
	fmt.Println("💾 Hippocampus Agent initialized.")
	return &HippocampusAgent{BaseAgent: NewAgent("Hippocampus")}
}

// ProcessTask handles memory-related tasks.
func (a *HippocampusAgent) ProcessTask(task *utils.Task) {
	fmt.Printf("💾 Hippocampus Agent managing memory task: %s\n", task.Data)

	// Store task in long-term memory (ChromaDB)
	a.storeMemory(task)
}

// storeMemory saves tasks in ChromaDB for long-term retrieval
func (a *HippocampusAgent) storeMemory(task *utils.Task) {
	// Ensure the collection is created or accessed in ChromaDB
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Hippocampus Agent failed to access ChromaDB:", err)
		return
	}

	// Convert text to vector using utility function
	vector := utils.GenerateVector(task.Data)

	// Store task as a vector in ChromaDB for long-term memory
	err = chroma.AddTaskToChroma(collectionID, utils.TaskVector{
		ID:       task.ID,
		TaskName: task.Data,
		Vector:   vector,
	})
	if err != nil {
		fmt.Println("❌ Failed to store task in long-term memory.")
		return
	}

	fmt.Println("✅ Task successfully stored in Hippocampus (Long-Term Memory).")
}

// RetrieveMemory fetches similar tasks from ChromaDB
func (a *HippocampusAgent) RetrieveMemory(taskData string) ([]utils.TaskVector, error) {
	// Ensure collection is accessible in ChromaDB
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		return nil, fmt.Errorf("failed to access ChromaDB")
	}

	// Convert taskData to a vector representation for the search
	queryVector := utils.GenerateVector(taskData)

	// Query ChromaDB to find the top 5 similar tasks
	tasks, err := chroma.SearchTaskInChroma(collectionID, queryVector, 5)
	if err != nil {
		return nil, fmt.Errorf("failed to query memory")
	}

	return tasks, nil
}

// LearnFromMemory updates the Hippocampus with learned experiences from previous tasks.
func (a *HippocampusAgent) LearnFromMemory(memory []utils.TaskVector) {
	fmt.Println("💾 Hippocampus Agent learning from past tasks...")

	// Learn from past tasks and store them in memory (ChromaDB)
	for _, task := range memory {
		// Store task in long-term memory (ChromaDB)
		a.storeMemory(&utils.Task{
			ID:   task.ID,
			Data: task.TaskName,
		})
	}

	// The Hippocampus is now updated by storing the learned tasks in ChromaDB
	fmt.Println("🧠 Memory updated with new learned tasks.")
}
