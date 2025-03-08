package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func ReportTaskStats() {
	file, err := os.Open("logs/tasks.log")
	if err != nil {
		fmt.Println("❌ Error opening log file:", err)
		return
	}
	defer file.Close()

	taskCounts := make(map[string]int)
	totalExecutionTime := make(map[string]time.Duration)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Task \d+ executed \| Priority: \d+ \| Duration: ([\d\.]+)(µ?s) \| Data: (.+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if len(matches) == 4 {
			durationStr := matches[1]
			unit := matches[2]
			taskName := matches[3]

			durationFloat, _ := strconv.ParseFloat(durationStr, 64)
			var duration time.Duration
			if unit == "µs" {
				duration = time.Duration(durationFloat) * time.Microsecond
			} else {
				duration = time.Duration(durationFloat) * time.Millisecond
			}

			taskCounts[taskName]++
			totalExecutionTime[taskName] += duration
		}
	}

	fmt.Println("\n📊 Task Execution Report:")
	fmt.Println("────────────────────────")
	for task, count := range taskCounts {
		avgTime := totalExecutionTime[task] / time.Duration(count)
		fmt.Printf("🔹 %s → Count: %d | Avg Execution Time: %v\n", task, count, avgTime)
	}
	fmt.Println("────────────────────────")
}
