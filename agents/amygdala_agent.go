package agents

import (
	"evolvai/utils"
	"fmt"
	"strings"
)

// AmygdalaAgent specializes in analyzing and prioritizing tasks.
type AmygdalaAgent struct {
    *BaseAgent
}

// NewAmygdalaAgent initializes the AI agent for Amygdala
func NewAmygdalaAgent() *AmygdalaAgent {
    return &AmygdalaAgent{
        BaseAgent: NewAgent("Amygdala"),
    }
}

// ProcessTask analyzes tasks and assigns priority.
func (a *AmygdalaAgent) ProcessTask(task *utils.Task) {
    fmt.Printf("🟠 AI Amygdala Agent processing: %s\n", task.Data)

    switch {
    case strings.Contains(strings.ToLower(task.Data), "error"),
        strings.Contains(strings.ToLower(task.Data), "critical"):
        task.Priority = utils.High
    case strings.Contains(strings.ToLower(task.Data), "update"),
        strings.Contains(strings.ToLower(task.Data), "memory"):
        task.Priority = utils.Medium
    case strings.Contains(strings.ToLower(task.Data), "self-awareness"),
        strings.Contains(strings.ToLower(task.Data), "pattern"):
        task.Priority = utils.Low
    default:
        task.Priority = utils.Low
    }

    fmt.Printf("🔹 AI Amygdala assigned priority %d to Task %d\n", task.Priority, task.ID)
}
