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
		Data: "System error detected", // This should be high priority
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.High {
		t.Errorf("Expected priority %d, but got %d for task %s", utils.High, task.Priority, task.Data)
	}

	// Test medium-priority task
	task = &utils.Task{
		ID:   2,
		Data: "Memory usage warning", // This should be medium priority
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.Medium {
		t.Errorf("Expected priority %d, but got %d for task %s", utils.Medium, task.Priority, task.Data)
	}

	// Test low-priority task
	task = &utils.Task{
		ID:   3,
		Data: "Pattern recognition triggered", // This should be low priority
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.Low {
		t.Errorf("Expected priority %d, but got %d for task %s", utils.Low, task.Priority, task.Data)
	}

	// Test unrecognized task (should be low priority)
	task = &utils.Task{
		ID:   4,
		Data: "Random task without any specific keyword", // This should default to low priority
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.Low {
		t.Errorf("Expected priority %d, but got %d for task %s", utils.Low, task.Priority, task.Data)
	}

	// Test repeated task (same task should get the same priority)
	task = &utils.Task{
		ID:   5,
		Data: "System error detected", // This should be recognized from memory and remain high priority
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.High {
		t.Errorf("Expected priority %d, but got %d for task %s (repeated)", utils.High, task.Priority, task.Data)
	}

	// Test a task with a new keyword, ensuring it's assigned the right priority
	task = &utils.Task{
		ID:   6,
		Data: "Urgent system update", // New high-priority task
	}
	amygdalaAgent.ProcessTask(task)
	if task.Priority != utils.High {
		t.Errorf("Expected priority %d, but got %d for task %s", utils.High, task.Priority, task.Data)
	}
}
