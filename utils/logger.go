package utils

import (
	"fmt"
	"time"
)

// LogInfo prints informational logs
func LogInfo(message string) {
	fmt.Printf("[%s] INFO: %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}
