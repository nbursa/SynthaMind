package cortex

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"evolvai/chroma"
	"evolvai/modules"
	"evolvai/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// short-term memory buffer
var recentEvents []string

// CortexBase runs autonomic processes like self-monitoring and system awareness
func CortexBase() {
	fmt.Println("🧠 CortexBase: Initializing AI self-awareness...")

	systemTicker := time.NewTicker(10 * time.Second)
	reviewTicker := time.NewTicker(60 * time.Second) // Every 60 seconds review memory

	for {
		select {
		case <-systemTicker.C:
			analyzeSystemState()
			storeSelfKnowledge()
		case <-reviewTicker.C:
			reviewRecentEvents()
		}
	}
}

// Analyze system state by fetching multiple system metrics
func analyzeSystemState() {
	// CPU usage
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

	// Memory usage
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("❌ Error getting memory usage:", err)
		return
	}
	memoryUsagePercent := memoryStats.UsedPercent

	// Disk usage
	diskUsage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("❌ Error getting disk usage:", err)
		return
	}
	diskAvailablePercent := 100.0 - diskUsage.UsedPercent

	// Network latency
	networkLatency := getNetworkLatency("8.8.8.8:53")

	// CPU temperature (mock)
	cpuTemperature := getTemperature()

	// Log system metrics
	event := fmt.Sprintf("CPU: %.2f%% | Memory: %.2f%% | Disk: %.2f%% | Latency: %.2fms | CPU Temp: %.2f°C",
		avgCPU, memoryUsagePercent, diskAvailablePercent, networkLatency, cpuTemperature)
	fmt.Println("⚙️  System Check:", event)

	// Remember event
	rememberEvent(event)

	// Reflex triggers
	if avgCPU > 90.0 || memoryUsagePercent > 90.0 {
		fmt.Println("🚨 WARNING: High system load! Activating survival mode...")
		survivalMechanism()
	}
}

// Check network latency by measuring round-trip time
func getNetworkLatency(server string) float64 {
	start := time.Now()
	conn, err := net.Dial("udp", server)
	if err != nil {
		fmt.Println("❌ Network latency error:", err)
		return 0.0
	}
	defer conn.Close()
	return time.Since(start).Seconds() * 1000 // ms
}

// Mock CPU temperature
func getTemperature() float64 {
	return 65.0 // Placeholder temperature
}

// survivalMechanism activated under high load
func survivalMechanism() {
	fmt.Println("🔄 CortexBase: Survival mode activated due to overload.")
	rememberEvent("Survival mode activated.")
}

// discoverEnvironment scans filesystem
func discoverEnvironment() []string {
	var discoveries []string
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("❌ Environment discovery error:", err)
		return discoveries
	}
	for _, entry := range entries {
		if entry.IsDir() {
			discoveries = append(discoveries, "📁 Directory: "+entry.Name())
		} else {
			switch {
			case strings.HasSuffix(entry.Name(), ".json"), strings.HasSuffix(entry.Name(), ".yaml"), strings.HasSuffix(entry.Name(), ".toml"):
				discoveries = append(discoveries, "🗃️ Config File: "+entry.Name())
			case strings.HasSuffix(entry.Name(), ".go"):
				discoveries = append(discoveries, "📜 Go Source: "+entry.Name())
			case strings.HasSuffix(entry.Name(), ".sh"):
				discoveries = append(discoveries, "📜 Shell Script: "+entry.Name())
			default:
				discoveries = append(discoveries, "📄 File: "+entry.Name())
			}
		}
	}
	return discoveries
}

func storeSelfKnowledge() {
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ ChromaDB collection error:", err)
		return
	}

	hostname, _ := os.Hostname()
	discoveries := discoverEnvironment()

	metadata := map[string]interface{}{
		"hostname":    hostname,
		"os":          runtime.GOOS,
		"cpu":         runtime.NumCPU(),
		"discoveries": discoveries,
		"timestamp":   time.Now().Format(time.RFC3339),
	}

	payload, _ := json.Marshal(metadata)

	task := utils.TaskVector{
		ID:       int(time.Now().Unix()),
		TaskName: "Self-awareness update",
		Vector:   modules.GenerateVector(string(payload)),
	}

	fmt.Println("📤 Storing self-awareness data in ChromaDB:", string(payload))

	err = chroma.AddTaskToChroma(collectionID, task)
	if err != nil {
		fmt.Println("❌ Failed to store data:", err)
	}

	// Remember event
	rememberEvent("Self-awareness data stored in ChromaDB.")
}

// rememberEvent saves recent events (short-term memory)
func rememberEvent(event string) {
	const maxEvents = 10
	if len(recentEvents) >= maxEvents {
		recentEvents = recentEvents[1:]
	}
	recentEvents = append(recentEvents, event)
}

// recallEvents retrieves recent events
func recallEvents() []string {
	return recentEvents
}

// reviewRecentEvents periodically reviews short-term memory
func reviewRecentEvents() {
	fmt.Println("🧠 Reviewing recent events in short-term memory...")
	events := recallEvents()
	for i, event := range events {
		fmt.Printf("📌 Event %d: %s\n", i+1, event)
	}
}
