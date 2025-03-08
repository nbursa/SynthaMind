package agents

import (
	"evolvai/chroma"
	"evolvai/utils"
	"fmt"
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

	// ✅ Store task in long-term memory
	a.storeMemory(task)
}

// storeMemory saves tasks in ChromaDB for long-term retrieval
func (a *HippocampusAgent) storeMemory(task *utils.Task) {
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Hippocampus Agent failed to access ChromaDB.")
		return
	}

	vector := utils.GenerateVector(task.Data) // ✅ Convert text to vector

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
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		return nil, fmt.Errorf("failed to access ChromaDB")
	}

	queryVector := utils.GenerateVector(taskData)
	tasks, err := chroma.SearchTaskInChroma(collectionID, queryVector, 5) // ✅ Retrieve top 5 matches
	if err != nil {
		return nil, fmt.Errorf("failed to query memory")
	}

	return tasks, nil
}
