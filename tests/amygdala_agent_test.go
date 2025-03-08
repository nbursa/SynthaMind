package tests

import (
	"evolvai/agents"
	"evolvai/utils"
	"testing"
)

func TestAmygdalaAgent_ProcessTask(t *testing.T) {
	// Setup
	amygdalaAgent := agents.NewAmygdalaAgent()

	// Define test cases
	tests := []struct {
		taskData       string
		expectedPriority utils.TaskPriority
	}{
		{"System error detected", utils.High},
		{"Memory usage warning", utils.Medium},
		{"Pattern recognition triggered", utils.Low},
		{"Urgent system update", utils.High},  // High priority based on "urgent"
		{"Random task without any specific keyword", utils.Low},
	}

	// Iterate through the test cases
	for _, test := range tests {
		t.Run(test.taskData, func(t *testing.T) {
			task := &utils.Task{
				ID:   1,
				Data: test.taskData,
			}

			// Process the task
			amygdalaAgent.ProcessTask(task)

			// Log the result for debugging purposes
			t.Logf("Task '%s' processed with priority %d", test.taskData, task.Priority)

			// Check if the expected priority matches
			if task.Priority != test.expectedPriority {
				t.Errorf("Expected priority %d, but got %d for task: %s", test.expectedPriority, task.Priority, test.taskData)
			} else {
				t.Logf("Task '%s' passed with expected priority: %d", test.taskData, test.expectedPriority)
			}
		})
	}
}
