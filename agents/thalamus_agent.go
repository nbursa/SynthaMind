package agents

import (
	"fmt"
	"strings"
	"synthamind/utils"
)

// ThalamusAgent filters sensory data and determines task importance.
type ThalamusAgent struct {
	*BaseAgent
	learnedKeywords map[string]int // Stores frequently occurring words from memory
}

// NewThalamusAgent initializes the Thalamus AI agent.
func NewThalamusAgent() *ThalamusAgent {
	fmt.Println("🔵 Thalamus Agent initialized.")
	return &ThalamusAgent{
		BaseAgent:       NewAgent("Thalamus"),
		learnedKeywords: make(map[string]int), // Stores dynamically learned words
	}
}

// ProcessTask filters and routes tasks appropriately.
func (a *ThalamusAgent) ProcessTask(task *utils.Task) {
	fmt.Printf("🔵 Thalamus Agent filtering task: %s\n", task.Data)

	// Check if task is important using dynamic learning
	if a.isImportant(task.Data) {
		task.Priority = utils.High // Assign high priority if task is important
		fmt.Println("✅ Task is important! Passing to Cortex...")
	} else {
		task.Priority = utils.Low // Assign low priority if task is not important
		fmt.Println("🚫 Task is NOT important. Discarding.")
	}
}

// Retrieve memory & update learned keywords
func (a *ThalamusAgent) LearnFromMemory(memory []utils.TaskVector) {
	fmt.Println("🧠 Thalamus Agent learning from past important tasks...")

	for _, task := range memory {
		words := strings.Fields(strings.ToLower(task.TaskName))
		for _, word := range words {
			a.learnedKeywords[word]++ // Track word occurrences
		}
	}

	fmt.Println("🔍 Updated learned keywords:", a.learnedKeywords)
}

// Check if a task is important using both static & learned rules
func (a *ThalamusAgent) isImportant(data string) bool {
	// Static keywords (hardcoded)
	staticKeywords := []string{"error", "critical", "update", "alert", "memory", "failure"}

	// Check if any static keyword is found
	for _, keyword := range staticKeywords {
		if strings.Contains(strings.ToLower(data), keyword) {
			fmt.Printf("✅ Matched static keyword: '%s' → Task is important!\n", keyword)
			return true
		}
	}

	// Check learned keywords (from past memory)
	words := strings.Fields(strings.ToLower(data))
	for _, word := range words {
		if a.learnedKeywords[word] > 2 { // If word appears frequently in memory
			fmt.Printf("✅ Matched learned keyword: '%s' → Task is important!\n", word)
			return true
		}
	}

	return false
}
