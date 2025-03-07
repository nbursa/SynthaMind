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
		go CortexProcess(task)
	} else {
		fmt.Println("🚫 Non-critical data detected, discarding...")
	}
}

func isImportant(data string) bool {
	importantKeywords := []string{
		"error", "warning", "alert", "update", "critical",
		"self-awareness", "discovery", "environment", "memory", "pattern",
	}

	// 🔍 Debug: Print each keyword check
	fmt.Printf("🔎 Checking if task is important: '%s'\n", data)

	for _, keyword := range importantKeywords {
		if strings.Contains(strings.ToLower(data), keyword) {
			fmt.Printf("✅ Matched keyword: '%s' → Task is important!\n", keyword)
			return true
		}
	}

	fmt.Println("❌ No important keywords found. Task is NOT important.")
	return false
}
