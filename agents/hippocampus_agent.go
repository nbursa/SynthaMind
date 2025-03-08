package agents

import (
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

// ✅ ProcessTask now accepts `utils.Task` instead of `string`
func (a *HippocampusAgent) ProcessTask(task *utils.Task) {
    fmt.Printf("💾 Hippocampus Agent managing memory task: %s\n", task.Data)
    // TODO: Implement memory management logic here.
}
