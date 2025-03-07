package cortex

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"evolvai/chroma"
	"evolvai/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
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

// Analyze system state by fetching multiple system metrics
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
	memoryUsagePercent := memoryStats.UsedPercent

	// Fetch available disk space
	diskUsage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("❌ Error getting disk usage:", err)
		return
	}
	diskAvailablePercent := 100.0 - diskUsage.UsedPercent

	// Check network latency to a common public DNS server (e.g., Google's 8.8.8.8)
	networkLatency := getNetworkLatency("8.8.8.8:53")

	// Mock CPU temperature
	cpuTemperature := getTemperature() // Placeholder for actual temperature reading

	fmt.Printf("⚙️  System Check: CPU: %.2f%% | Memory: %.2f%% | Disk Available: %.2f%% | Network Latency: %.2fms | CPU Temp: %.2f°C\n",
		avgCPU, memoryUsagePercent, diskAvailablePercent, networkLatency, cpuTemperature)

	if avgCPU > 90.0 || memoryUsagePercent > 90.0 {
		fmt.Println("🚨 WARNING: High system load! Taking preventive action...")
		survivalMechanism()
	}
}

// Check network latency by measuring round-trip time
func getNetworkLatency(server string) float64 {
	start := time.Now()
	conn, err := net.Dial("udp", server)
	if err != nil {
		fmt.Println("❌ Error checking network latency:", err)
		return 0.0
	}
	defer conn.Close()

	latency := time.Since(start).Seconds() * 1000 // Convert to milliseconds
	return latency
}

// Mock function for CPU temperature
func getTemperature() float64 {
	// Placeholder: in a real implementation, read from a thermal sensor or system file
	return 65.0 // Example: 65°C
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
		"os":       runtime.GOOS,
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
