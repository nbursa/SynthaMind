package modules

import (
	"evolvai/utils"
	"fmt"
)

// ThalamusFilter processes raw input
func ThalamusFilter(task utils.Task) {
	fmt.Printf("🔵 Thalamus Filtering Task %d: %s\n", task.ID, task.Data)

	if len(task.Data) > 10 {
		fmt.Println("✅ Important data detected, passing to Cortex...")
		CortexProcess(task) // Send data to Cortex
	} else {
		fmt.Println("❌ Low-priority data, ignoring.")
	}
}