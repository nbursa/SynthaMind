package agents

import (
	"evolvai/utils"
	"fmt"
	"strings"
)

// AmygdalaAgent specializes in analyzing and prioritizing tasks.
type AmygdalaAgent struct {
    *BaseAgent
    memory map[string]utils.TaskPriority // ✅ Store past task data
}

// NewAmygdalaAgent initializes the AI agent for Amygdala
func NewAmygdalaAgent() *AmygdalaAgent {
    return &AmygdalaAgent{
        BaseAgent: NewAgent("Amygdala"),
        memory:    make(map[string]utils.TaskPriority), // ✅ Initialize memory storage
    }
}

// ProcessTask analyzes tasks and assigns priority.
func (a *AmygdalaAgent) ProcessTask(task *utils.Task) {
    fmt.Printf("🟠 AI Amygdala Agent processing: %s\n", task.Data)

    // ✅ Check if task was seen before
    if priority, exists := a.memory[task.Data]; exists {
        task.Priority = priority // ✅ Use learned priority
        fmt.Printf("🔁 AI Amygdala RECOGNIZED task, setting previous priority: %d\n", priority)
    } else {
        // ✅ If new task, analyze it
        a.analyzeTask(task)
        // ✅ Store decision in memory
        a.memory[task.Data] = task.Priority
    }
}

// ✅ Task Analysis Logic (now adaptive)
func (a *AmygdalaAgent) analyzeTask(task *utils.Task) {
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
