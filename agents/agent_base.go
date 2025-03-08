package agents

import "fmt"

// BaseAgent provides a structure for AI agents controlling different modules
type BaseAgent struct {
	Name string
}

// NewAgent initializes a new AI agent
func NewAgent(name string) *BaseAgent {
	return &BaseAgent{Name: name}
}

// ProcessTask allows an agent to process a task
func (a *BaseAgent) ProcessTask(task string) {
	fmt.Printf("🤖 AI Agent [%s] processing task: %s\n", a.Name, task)
}
