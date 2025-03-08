package agents

import (
	"evolvai/utils"
	"fmt"
	"math"
	"strings"
)

// CortexAgent is responsible for higher-level cognitive processing.
type CortexAgent struct {
	*BaseAgent
	Hippocampus *HippocampusAgent // Inject HippocampusAgent
}

// NewCortexAgent initializes the Cortex AI agent.
func NewCortexAgent(hippocampus *HippocampusAgent) *CortexAgent {
	fmt.Println("🧠 Cortex Agent initialized.")
	return &CortexAgent{
		BaseAgent:   NewAgent("Cortex"),
		Hippocampus: hippocampus, // Store reference
	}
}

// ProcessTask handles cognitive processing.
func (a *CortexAgent) ProcessTask(task *utils.Task) {
	fmt.Printf("🧠 Cortex Agent processing task: %s\n", task.Data)

	// Retrieve past similar tasks from Hippocampus
	similarTasks, err := a.Hippocampus.RetrieveMemory(task.Data)
	if err != nil || len(similarTasks) == 0 {
		fmt.Println("❌ No past memory found. Processing as a new task.")
		a.processNewTask(task)
		return
	}

	// If similar tasks exist, analyze them
	a.analyzeTask(task, similarTasks)
}

// Handle a completely new task with no memory
func (a *CortexAgent) processNewTask(task *utils.Task) {
	fmt.Printf("🧠 New task detected: %s\n", task.Data)
	fmt.Println("🤔 No past cases found. Handling task using default rules.")
	// Implement rule-based decision logic
}

// Compare new task with past memory & decide action
func (a *CortexAgent) analyzeTask(task *utils.Task, pastTasks []utils.TaskVector) {
	fmt.Println("🔍 Cortex analyzing task based on memory...")

	// Rank past tasks based on similarity (using cosine similarity, for example)
	bestMatch := a.findBestMatch(task, pastTasks)

	fmt.Printf("🧠 Cortex found a similar past task: %s\n", bestMatch.TaskName)

	// Determine next action based on past experience
	if bestMatch.TaskName == task.Data {
		fmt.Println("✅ Exact past task match! Following previous solution.")
	} else {
		fmt.Println("🤔 Partial match found. Adapting approach.")
		a.adaptTaskResponse(task, bestMatch)
	}
}

// Compare task with past tasks using some similarity function
func (a *CortexAgent) findBestMatch(task *utils.Task, pastTasks []utils.TaskVector) utils.TaskVector {
	var bestMatch utils.TaskVector
	maxSimilarity := -math.MaxFloat64

	// Use cosine similarity or other heuristics to rank tasks
	for _, past := range pastTasks {
		similarity := a.calculateSimilarity(task.Data, past.TaskName)
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			bestMatch = past
		}
	}

	return bestMatch
}

// Calculate similarity between two strings (e.g., cosine similarity)
func (a *CortexAgent) calculateSimilarity(taskData, taskName string) float64 {
	// Example: simple string length-based similarity (can be improved)
	return float64(len(taskData)) / float64(len(taskName))
}

// Modify the response based on learned patterns
func (a *CortexAgent) adaptTaskResponse(task *utils.Task, bestMatch utils.TaskVector) {
	fmt.Println("🔄 Cortex is modifying the response based on learned patterns...")

	// Example modification (expand logic later)
	if strings.Contains(bestMatch.TaskName, "error") && !strings.Contains(task.Data, "error") {
		fmt.Println("🛠️ Similar issue but NOT an error. Adjusting priority lower.")
		task.Priority = utils.Medium
	} else {
		fmt.Println("🔄 No direct match found. Using best available past response.")
	}

	fmt.Printf("🧠 Adapted Task %d → New Priority: %d\n", task.ID, task.Priority)
}
