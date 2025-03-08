package tests

import (
	"evolvai/agents"
	"evolvai/utils"
	"testing"
)

func TestAmygdalaAgent_ProcessTask(t *testing.T) {
	// Setup
	amygdalaAgent := agents.NewAmygdalaAgent()

	// Test high-priority task
	task := &utils.Task{
		ID:   1,
		Data: "System error detected",
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.High {
		t.Errorf("Expected priority %d, but got %d", utils.High, task.Priority)
	}

	// Test medium-priority task
	task = &utils.Task{
		ID:   2,
		Data: "Memory usage warning",
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.Medium {
		t.Errorf("Expected priority %d, but got %d", utils.Medium, task.Priority)
	}

	// Test low-priority task
	task = &utils.Task{
		ID:   3,
		Data: "Pattern recognition triggered",
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.Low {
		t.Errorf("Expected priority %d, but got %d", utils.Low, task.Priority)
	}
}
