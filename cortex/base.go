package cortex

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"

	"evolvai/chroma"
	"evolvai/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// CortexBase runs autonomic processes like self-monitoring and system awareness
func CortexBase() {
	fmt.Println("🧠 CortexBase: Initializing AI self-awareness...")

	for {
		analyzeSystemState()
		storeSelfKnowledge()
		time.Sleep(10 * time.Second) // Periodic check every 10s
	}
}

// // Analyze system state (memory, CPU, files, APIs)
// func analyzeSystemState() {
// 	var memStats runtime.MemStats
// 	runtime.ReadMemStats(&memStats)

// 	cpuUsage := getCPUUsage()
// 	memoryUsage := memStats.Alloc / 1024 / 1024 // Convert bytes to MB

// 	fmt.Printf("⚙️  System Check: CPU: %.2f%% | Memory: %dMB\n", cpuUsage, memoryUsage)

// 	if cpuUsage > 90.0 || memoryUsage > 500 {
// 		fmt.Println("🚨 WARNING: High system load! Taking preventive action...")
// 		survivalMechanism()
// 	}
// }

// // Mock CPU usage retrieval
// func getCPUUsage() float64 {
// 	return 30.0 // TODO: Implement actual CPU monitoring
// }

func analyzeSystemState() {
    // Fetch CPU usage
    cpuUsage, err := cpu.Percent(0, false)
    if err != nil {
        fmt.Println("❌ Error getting CPU usage:", err)
        return
    }
    totalCPU := 0.0
    for _, usage := range cpuUsage {
        totalCPU += usage
    }
    avgCPU := totalCPU / float64(len(cpuUsage))

    // Fetch memory usage
    memoryStats, err := mem.VirtualMemory()
    if err != nil {
        fmt.Println("❌ Error getting memory usage:", err)
        return
    }
    // memoryUsage := memoryStats.Used / 1024 / 1024 // Convert bytes to MB
	memoryUsagePercent := memoryStats.UsedPercent


    fmt.Printf("⚙️  System Check: CPU: %.2f%% | Memory: %.2f%%\n", avgCPU, memoryUsagePercent)

	if avgCPU > 90.0 || memoryUsagePercent > 90.0 {
		fmt.Println("🚨 WARNING: High system load! Taking preventive action...")
		survivalMechanism()
	}
}

// If the system is overloaded, take preventive actions
func survivalMechanism() {
	fmt.Println("🔄 CortexBase: Activating survival mode...")
	// Future: Reduce processing load, free up memory, or alert admin
}

// Store what CortexBase learns about its environment in ChromaDB
func storeSelfKnowledge() {
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Failed to ensure ChromaDB collection:", err)
		return
	}

	hostname, _ := os.Hostname()
	metadata := map[string]interface{}{
		"hostname": hostname,
		"os":      runtime.GOOS,
		"cpu":      runtime.NumCPU(),
		"timestamp": time.Now().Format(time.RFC3339),
	}

	payload, _ := json.Marshal(metadata)
	task := utils.TaskVector{
		ID:       int(time.Now().Unix()),
		TaskName: "Self-awareness update",
		Vector:   []float32{1.0, 0.5, 0.2}, // Placeholder
	}

	fmt.Println("📤 Storing self-awareness update in ChromaDB:", string(payload))
	err = chroma.AddTaskToChroma(collectionID, task)
	if err != nil {
		fmt.Println("❌ Failed to store self-awareness data:", err)
	}
}
