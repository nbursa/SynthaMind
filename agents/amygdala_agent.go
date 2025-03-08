package agents

import (
	"evolvai/utils"
	"fmt"
)

// AmygdalaAgent specializes in emotional response and priority processing
type AmygdalaAgent struct {
	*BaseAgent
	AnalyzeFunc func(*utils.Task) // ✅ Function parameter to avoid direct dependency
}

// NewAmygdalaAgent initializes the AI agent for Amygdala
func NewAmygdalaAgent(analyzeFunc func(*utils.Task)) *AmygdalaAgent {
	return &AmygdalaAgent{
		BaseAgent:   NewAgent("Amygdala"), // ✅ Initialize base agent
		AnalyzeFunc: analyzeFunc,
	}
}

// ProcessTask prioritizes tasks like a real amygdala
func (a *AmygdalaAgent) ProcessTask(task *utils.Task) {
	fmt.Printf("🟠 AI Amygdala Agent processing: %s\n", task.Data)
	a.AnalyzeFunc(task) // ✅ Uses the injected function instead of direct module dependency
}
