package modules

import (
	"evolvai/utils"
	"fmt"
	"strings"
)

// ThalamusFilter acts as a sensory gate, filtering and prioritizing incoming tasks.
func ThalamusFilter(task utils.Task) {
	fmt.Printf("🔵 Thalamus Filtering Task %d: %s\n", task.ID, task.Data)

	if isImportant(task.Data) {
		fmt.Println("✅ Important data detected, passing to Cortex...")
		go CortexProcess(task) // Corrected reference
	} else {
		fmt.Println("🚫 Non-critical data detected, discarding...")
	}
}

// isImportant evaluates task importance based on predefined keywords
func isImportant(data string) bool {
	importantKeywords := []string{
		"error", "warning", "alert", "update", "critical",
		"self-awareness", "discovery", "environment", "memory", "pattern",
	}

	for _, keyword := range importantKeywords {
		if strings.Contains(strings.ToLower(data), keyword) {
			return true
		}
	}

	return false
}
