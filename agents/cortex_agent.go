package agents

import (
	"evolvai/utils"
	"fmt"
)

// CortexAgent is responsible for higher-level cognitive processing.
type CortexAgent struct {
	*BaseAgent
	Hippocampus *HippocampusAgent // ✅ Inject HippocampusAgent
}

// NewCortexAgent initializes the Cortex AI agent.
func NewCortexAgent(hippocampus *HippocampusAgent) *CortexAgent {
	fmt.Println("🧠 Cortex Agent initialized.")
	return &CortexAgent{
		BaseAgent:   NewAgent("Cortex"),
		Hippocampus: hippocampus, // ✅ Store reference
	}
}

// ProcessTask handles cognitive processing.
func (a *CortexAgent) ProcessTask(task *utils.Task) {
	fmt.Printf("🧠 Cortex Agent processing task: %s\n", task.Data)

	// ✅ Retrieve past similar tasks from Hippocampus
	similarTasks, err := a.Hippocampus.RetrieveMemory(task.Data)
	if err != nil || len(similarTasks) == 0 {
		fmt.Println("❌ No past memory found. Processing as a new task.")
		a.processNewTask(task)
		return
	}

	// ✅ If similar tasks exist, analyze them
	a.analyzeTask(task, similarTasks)
}

// ✅ Handle a completely new task with no memory
func (a *CortexAgent) processNewTask(task *utils.Task) {
	fmt.Printf("🧠 New task detected: %s\n", task.Data)
	fmt.Println("🤔 No past cases found. Handling task using default rules.")
	// TODO: Implement default rule-based decision logic
}

// ✅ Compare new task with past memory & decide action
func (a *CortexAgent) analyzeTask(task *utils.Task, pastTasks []utils.TaskVector) {
	fmt.Println("🔍 Cortex analyzing task based on memory...")

	// ✅ Find the most relevant past task
	bestMatch := pastTasks[0] // Assume first is most relevant
	for _, past := range pastTasks {
		if len(past.TaskName) > len(bestMatch.TaskName) { // Simple heuristic: longer match
			bestMatch = past
		}
	}

	fmt.Printf("🧠 Cortex found a similar past task: %s\n", bestMatch.TaskName)

	// ✅ Determine next action based on past experience
	if bestMatch.TaskName == task.Data {
		fmt.Println("✅ Exact past task match! Following previous solution.")
		// TODO: Implement logic to follow past resolution strategy
	} else {
		fmt.Println("🤔 Partial match found. Adapting approach.")
		// TODO: Implement logic to modify the response dynamically
	}
}
