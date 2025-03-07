package modules

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"evolvai/chroma"
	"evolvai/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Short-term memory buffer
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

// CortexProcess processes tasks and integrates memory storage
func CortexProcess(task utils.Task) {
	fmt.Printf("🧠 Cortex Processing Task %d: %s\n", task.ID, task.Data)

	// 🔍 Step 1: Check if task exists in Hippocampus (memory)
	memory := HippocampusRetrieve(task)
	if memory != nil {
		fmt.Println("🔁 Using past memory instead of reprocessing.")
		return
	}

	// 🔹 Step 2: Convert task text to a vector
	vector := GenerateVector(task.Data)

	// 🔹 Step 3: Ensure ChromaDB collection exists
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Skipping task processing due to ChromaDB unavailability.")
		return
	}

	// 🔹 Step 4: Store new knowledge
	err = chroma.AddTaskToChroma(collectionID, utils.TaskVector{
		ID:       task.ID,
		TaskName: task.Data,
		Vector:   vector,
	})
	if err != nil {
		fmt.Println("❌ Failed to store task in ChromaDB.")
		return
	}

	// 🔹 Step 5: Store in Hippocampus memory
	HippocampusStore(task)
	fmt.Println("✅ New knowledge stored in Hippocampus.")
}

// Analyze system state by fetching multiple system metrics
func analyzeSystemState() {
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

	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("❌ Error getting memory usage:", err)
		return
	}
	memoryUsagePercent := memoryStats.UsedPercent

	diskUsage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("❌ Error getting disk usage:", err)
		return
	}
	diskAvailablePercent := 100.0 - diskUsage.UsedPercent

	networkLatency := getNetworkLatency("8.8.8.8:53")

	cpuTemperature := getTemperature()

	event := fmt.Sprintf("CPU: %.2f%% | Memory: %.2f%% | Disk: %.2f%% | Latency: %.2fms | CPU Temp: %.2f°C",
		avgCPU, memoryUsagePercent, diskAvailablePercent, networkLatency, cpuTemperature)
	fmt.Println("⚙️  System Check:", event)

	rememberEvent(event)

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
	return time.Since(start).Seconds() * 1000
}

// Mock CPU temperature
func getTemperature() float64 {
	return 65.0 // Placeholder temperature
}

// SurvivalMechanism activated under high load
func survivalMechanism() {
	fmt.Println("🔄 CortexBase: Survival mode activated due to overload.")
	rememberEvent("Survival mode activated.")
}

// Discover available system components and store them in memory
func discoverSystemComponents() []string {
	var components []string

	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			components = append(components, fmt.Sprintf("🌐 Network Interface: %s", iface.Name))
		}
	}

	components = append(components, fmt.Sprintf("🖥️ CPU Cores: %d", runtime.NumCPU()))

	memoryStats, err := mem.VirtualMemory()
	if err == nil {
		components = append(components, fmt.Sprintf("💾 Memory Available: %.2f GB", float64(memoryStats.Available)/1e9))
	}

	diskUsage, err := disk.Usage("/")
	if err == nil {
		components = append(components, fmt.Sprintf("📀 Disk Space Available: %.2f GB", float64(diskUsage.Free)/1e9))
	}

	entries, err := os.ReadDir(".")
	if err == nil {
		for _, entry := range entries {
			if strings.HasSuffix(entry.Name(), ".json") || strings.HasSuffix(entry.Name(), ".yaml") {
				components = append(components, fmt.Sprintf("⚙️ Config File: %s", entry.Name()))
			}
		}
	}

	fmt.Println("🔍 Self-Discovery Completed. Found Components:", components)
	return components
}

// Store self-awareness knowledge in ChromaDB
func storeSelfKnowledge() {
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ ChromaDB collection error:", err)
		return
	}

	hostname, _ := os.Hostname()
	components := discoverSystemComponents()

	metadata := map[string]interface{}{
		"hostname":    hostname,
		"os":          runtime.GOOS,
		"cpu":         runtime.NumCPU(),
		"components":  components,
		"timestamp":   time.Now().Format(time.RFC3339),
	}

	payload, _ := json.Marshal(metadata)

	task := utils.TaskVector{
		ID:       int(time.Now().Unix()),
		TaskName: "Self-awareness update",
		Vector:   GenerateVector(string(payload)),
	}

	fmt.Println("📤 Storing self-awareness data in ChromaDB:", string(payload))
	err = chroma.AddTaskToChroma(collectionID, task)
	if err != nil {
		fmt.Println("❌ Failed to store self-awareness data:", err)
	}
}

// GenerateVector converts text into a hashed float32 vector
func GenerateVector(data string) []float32 {
	hash := sha256.Sum256([]byte(data))
	vector := make([]float32, 3)

	for i := 0; i < 3; i++ {
		vector[i] = float32(binary.BigEndian.Uint32(hash[i*8 : (i+1)*8])) / 1e9
	}
	return vector
}

// WaitForChromaDB actively checks if ChromaDB is ready
func WaitForChromaDB() bool {
	fmt.Println("⏳ Waiting for ChromaDB to be available...")

	for i := 0; i < 5; i++ {
		resp, err := http.Get("http://127.0.0.1:8000/api/v1/heartbeat")
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("✅ ChromaDB is ready!")
			return true
		}
		fmt.Println("⚠️ ChromaDB not available, retrying...")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("❌ ChromaDB is still unavailable after retries.")
	return false
}

// RememberEvent saves recent events (short-term memory)
func rememberEvent(event string) {
	const maxEvents = 10
	if len(recentEvents) >= maxEvents {
		recentEvents = recentEvents[1:]
	}
	recentEvents = append(recentEvents, event)
}

// RecallEvents retrieves recent events
func recallEvents() []string {
	return recentEvents
}

// ReviewRecentEvents periodically reviews short-term memory
func reviewRecentEvents() {
	fmt.Println("🧠 Reviewing recent events in short-term memory...")
	events := recallEvents()
	for i, event := range events {
		fmt.Printf("📌 Event %d: %s\n", i+1, event)
	}
}
