package modules

import (
	"evolvai/utils"
	"fmt"
	"strings"
)

// AmygdalaAnalyze assigns priority levels to tasks
func AmygdalaAnalyze(task *utils.Task) {
	fmt.Printf("🟠 Amygdala Processing Task %d: %s\n", task.ID, task.Data)

	switch {
	case strings.Contains(strings.ToLower(task.Data), "error"),
		strings.Contains(strings.ToLower(task.Data), "critical"):
		task.Priority = utils.High // ✅ Use `utils.High`
	case strings.Contains(strings.ToLower(task.Data), "update"),
		strings.Contains(strings.ToLower(task.Data), "memory"):
		task.Priority = utils.Medium // ✅ Use `utils.Medium`
	case strings.Contains(strings.ToLower(task.Data), "self-awareness"),
		strings.Contains(strings.ToLower(task.Data), "pattern"):
		task.Priority = utils.Low // ✅ Use `utils.Low`
	default:
		task.Priority = utils.Low // ✅ Use `utils.Low`
	}

	fmt.Printf("🔹 Task %d assigned priority: %d\n", task.ID, task.Priority)
}
