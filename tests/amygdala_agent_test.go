// Package tests contains unit tests for SynthaMind AI agents.
package tests

import (
	"synthamind/agents"
	"synthamind/utils"
	"testing"
)

// TestAmygdalaAgent_ProcessTask verifies that tasks are assigned correct priorities.
func TestAmygdalaAgent_ProcessTask(t *testing.T) {
	// 🏗 Setup
	amygdalaAgent := agents.NewAmygdalaAgent()

	// 📌 Define test cases for different types of tasks
	tests := []struct {
		name            string
		taskData        string
		expectedPriority utils.TaskPriority
	}{
		{"Critical system error", "System error detected", utils.High},
		{"Medium priority warning", "Memory usage warning", utils.Medium},
		{"Low priority pattern recognition", "Pattern recognition triggered", utils.Low},
		{"High priority urgent update", "Urgent system update", utils.High},
		{"Neutral task with no keywords", "Random task without any specific keyword", utils.Low},
		{"Empty task data", "", utils.Low}, // ✅ Edge case: Empty task should default to Low priority.
	}

	// 🔄 Iterate through test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := &utils.Task{
				ID:   1,
				Data: test.taskData,
			}

			// 🚀 Process the task
			amygdalaAgent.ProcessTask(task)

			// 📝 Log the result for debugging purposes
			t.Logf("Processing task: '%s' → Assigned Priority: %d", test.taskData, task.Priority)

			// ✅ Verify priority matches expected value
			if task.Priority != test.expectedPriority {
				t.Errorf("❌ Expected priority %d, but got %d for task: %s",
					test.expectedPriority, task.Priority, test.taskData)
			} else {
				t.Logf("✅ Task '%s' passed with expected priority: %d", test.taskData, test.expectedPriority)
			}
		})
	}
}
