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
		// No memory found, process as new task
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

    taskDataNormalized := strings.ToLower(strings.TrimSpace(task.Data)) // Normalize the task data

    // Tasks with "critical", "error", or "failure" should get high priority
    if strings.Contains(taskDataNormalized, "critical") || 
       strings.Contains(taskDataNormalized, "error") || 
       strings.Contains(taskDataNormalized, "failure") {
        task.Priority = utils.High
        fmt.Println("🛠️ Task is critical! Assigning high priority.")
    } else if strings.Contains(taskDataNormalized, "memory") || 
            strings.Contains(taskDataNormalized, "update") {
        // Tasks with "memory" or "update" should get medium priority
        task.Priority = utils.Medium
        fmt.Println("🟠 Task is of medium priority (memory, update).")
    } else {
        task.Priority = utils.Medium
        fmt.Println("🟠 Task is of medium priority.")
    }
}

// Compare new task with past memory & decide action
func (a *CortexAgent) analyzeTask(task *utils.Task, pastTasks []utils.TaskVector) {
	fmt.Println("🔍 Cortex analyzing task based on memory...")

	// Rank past tasks based on similarity (using cosine similarity, for example)
	bestMatch := a.findBestMatch(task, pastTasks)

	fmt.Printf("🧠 Cortex found a similar past task: %s\n", bestMatch.TaskName)

	// Determine next action based on past experience
	if bestMatch.TaskName == task.Data {
		// Exact match: retain high priority from past task
		fmt.Println("✅ Exact past task match! Following previous solution.")
		task.Priority = utils.High
	} else {
		// Partial match: adjust based on learned patterns
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

	// If no good match is found, return the default task with medium priority
	if maxSimilarity < 0.5 { // Adjust the threshold as needed
		fmt.Println("🤔 No good match found, adapting new task behavior.")
		task.Priority = utils.Medium
	} else {
		fmt.Println("✅ Found best match. Following past solution.")
	}

	return bestMatch
}

// Calculate similarity between two strings (e.g., cosine similarity)
func (a *CortexAgent) calculateSimilarity(taskData, taskName string) float64 {
	// Example: simple string length-based similarity (can be improved)
	taskDataWords := strings.Fields(strings.ToLower(taskData))
	taskNameWords := strings.Fields(strings.ToLower(taskName))

	// Create vectors for taskData and taskName (counting word frequencies)
	vectorA := make(map[string]int)
	vectorB := make(map[string]int)

	for _, word := range taskDataWords {
		vectorA[word]++
	}
	for _, word := range taskNameWords {
		vectorB[word]++
	}

	// Compute cosine similarity
	return cosineSimilarity(vectorA, vectorB)
}

// Calculate cosine similarity between two word vectors
func cosineSimilarity(vecA, vecB map[string]int) float64 {
	var dotProduct, magA, magB float64

	// Compute dot product and magnitudes
	for word, countA := range vecA {
		if countB, exists := vecB[word]; exists {
			dotProduct += float64(countA * countB)
		}
		magA += float64(countA * countA)
	}
	for _, countB := range vecB {
		magB += float64(countB * countB)
	}

	if magA == 0 || magB == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(magA) * math.Sqrt(magB))
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
