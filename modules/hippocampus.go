package modules

import (
	"evolvai/utils"
	"fmt"
)

// HippocampusStore saves knowledge
func HippocampusStore(task utils.Task) {
	fmt.Printf("💾 Storing Task %d in memory: %s\n", task.ID, task.Data)
}
