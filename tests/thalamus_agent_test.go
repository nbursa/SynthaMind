// Package tests contains unit tests for SynthaMind AI agents.
package tests

import (
	"evolvai/agents"
	"evolvai/utils"
	"testing"
)

// TestThalamusAgent_ProcessTask ensures important tasks are passed to Cortex, while unimportant ones are discarded.
func TestThalamusAgent_ProcessTask(t *testing.T) {
	// 🏗 Setup
	thalamusAgent := agents.NewThalamusAgent()

	// 📌 Define test cases for task filtering
	tests := []struct {
		name           string
		taskData       string
		expectedOutput string
		expectedPriority utils.TaskPriority
	}{
		{"Critical system error", "System error detected", "Task is important! Passing to Cortex...", utils.High},
		{"Medium severity warning", "Memory usage warning", "Task is important! Passing to Cortex...", utils.High},
		{"High priority urgent update", "Urgent system update", "Task is important! Passing to Cortex...", utils.High},
		{"Unimportant pattern recognition", "Pattern recognition triggered", "Task is NOT important. Discarding.", utils.Low},
		{"Neutral task without keywords", "Random task without any specific keyword", "Task is NOT important. Discarding.", utils.Low},
		{"Empty task data", "", "Task is NOT important. Discarding.", utils.Low}, // ✅ Edge case: Empty input should be discarded.
	}

	// 🔄 Iterate through test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := &utils.Task{
				ID:   1,
				Data: test.taskData,
			}

			// 🚀 Process the task
			thalamusAgent.ProcessTask(task)

			// 📝 Log task processing result
			t.Logf("Processing task: '%s'", test.taskData)

			// ✅ Validate task priority and expected handling
			if test.expectedOutput == "Task is important! Passing to Cortex..." && task.Priority != utils.High {
				t.Errorf("❌ Task '%s' expected to be High priority but was %d",
					test.taskData, task.Priority)
			} else if test.expectedOutput == "Task is NOT important. Discarding." && task.Priority != utils.Low {
				t.Errorf("❌ Task '%s' expected to be discarded but was assigned priority %d",
					test.taskData, task.Priority)
			} else {
				t.Logf("✅ Task '%s' correctly processed: %s", test.taskData, test.expectedOutput)
			}
		})
	}
}
