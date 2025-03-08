package tests

import (
	"evolvai/agents"
	"evolvai/utils"
	"testing"
)

func TestThalamusAgent_ProcessTask(t *testing.T) {
	// Setup
	thalamusAgent := agents.NewThalamusAgent()

	// Define test cases
	tests := []struct {
		taskData       string
		expectedOutput string
	}{
		{"System error detected", "Task is important! Passing to Cortex..."},
		{"Memory usage warning", "Task is important! Passing to Cortex..."},
		{"Urgent system update", "Task is important! Passing to Cortex..."},
		{"Pattern recognition triggered", "Task is NOT important. Discarding."},
		{"Random task without any specific keyword", "Task is NOT important. Discarding."},
	}

	// Iterate through the test cases
	for _, test := range tests {
		t.Run(test.taskData, func(t *testing.T) {
			task := &utils.Task{
				ID:   1,
				Data: test.taskData,
			}

			// Process the task
			thalamusAgent.ProcessTask(task)

			// Log the result for debugging purposes
			t.Logf("Task '%s' processed", test.taskData)

			// Check if the task is processed correctly
			if test.expectedOutput == "Task is important! Passing to Cortex..." && task.Priority == utils.High {
				t.Logf("Task '%s' is correctly marked as important and passed to Cortex", test.taskData)
			} else if test.expectedOutput == "Task is NOT important. Discarding." && task.Priority == utils.Low {
				t.Logf("Task '%s' correctly discarded", test.taskData)
			} else {
				t.Errorf("Task '%s' failed: Expected output: %s, but got: %d", test.taskData, test.expectedOutput, task.Priority)
			}
		})
	}
}
