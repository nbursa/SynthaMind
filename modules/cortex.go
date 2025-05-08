package modules

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"synthamind/chroma"
	"synthamind/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Short-term memory buffer for events
var recentEvents []string

// Memory storage for detected components
var detectedComponents = make(map[string]bool)

// CortexBase initializes AI self-awareness and system monitoring
func CortexBase() {
	fmt.Println("🧠 CortexBase: Initializing AI self-awareness...")

	//  Force self-discovery at startup
	selfDiscover()

	systemTicker := time.NewTicker(10 * time.Second)
	reviewTicker := time.NewTicker(60 * time.Second)    // Every 60 seconds review memory
	discoveryTicker := time.NewTicker(30 * time.Second) // Every 30s check system components

	for {
		select {
		case <-systemTicker.C:
			analyzeSystemState()
			storeSelfKnowledge()
		case <-discoveryTicker.C:
			selfDiscover()
		case <-reviewTicker.C:
			reviewRecentEvents()
		}
	}
}

// CortexProcess handles important tasks passed from ThalamusFilter
func CortexProcess(task utils.Task) {
	fmt.Printf("🧠 Cortex Processing Task %d: %s\n", task.ID, task.Data)

	// Check if task already exists in memory
	memory := HippocampusRetrieve(task)
	if memory != nil {
		fmt.Println("🔁 Using past memory instead of reprocessing.")
		return
	}

	// Convert task data to vector
	vector := GenerateVector(task.Data)

	// Ensure ChromaDB collection exists
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ Skipping task processing due to ChromaDB unavailability.")
		return
	}

	// Store task in ChromaDB
	err = chroma.AddTaskToChroma(collectionID, utils.TaskVector{
		ID:       task.ID,
		TaskName: task.Data,
		Vector:   vector,
	})
	if err != nil {
		fmt.Println("❌ Failed to store task in ChromaDB.")
		return
	}

	// Store task in memory
	HippocampusStore(task)
	fmt.Println("✅ New knowledge stored in Hippocampus.")

	// NEW: Execute task action
	Executor(task)
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
}

// Store system state knowledge in ChromaDB
func storeSelfKnowledge() {
	collectionID, err := chroma.EnsureChromaCollection()
	if err != nil {
		fmt.Println("❌ ChromaDB collection error:", err)
		return
	}

	hostname, _ := os.Hostname()
	discoveries := selfDiscover()

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
		Vector:   GenerateVector(string(payload)), // Ensure proper reference
	}

	fmt.Println("📤 Storing self-awareness data in ChromaDB:", string(payload))
	err = chroma.AddTaskToChroma(collectionID, task)
	if err != nil {
		fmt.Println("❌ Failed to store self-awareness data:", err)
	}
}

// Self-discovery of available system components
func selfDiscover() []string {
	fmt.Println("🔍 Cortex: Performing Self-Discovery...")

	components := []string{}

	// Check available network interfaces
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			component := fmt.Sprintf("🌐 Network Interface: %s", iface.Name)
			if !detectedComponents[component] {
				components = append(components, component)
				detectedComponents[component] = true
				fmt.Println("✅ Detected:", component) // 🔹 Debug Log
			}
		}
	}

	// Detect CPU cores
	cpuInfo := fmt.Sprintf("🖥️ CPU Cores: %d", runtime.NumCPU())
	if !detectedComponents[cpuInfo] {
		components = append(components, cpuInfo)
		detectedComponents[cpuInfo] = true
		fmt.Println("✅ Detected:", cpuInfo) // 🔹 Debug Log
	}

	// Detect memory
	memoryStats, err := mem.VirtualMemory()
	if err == nil {
		memInfo := fmt.Sprintf("💾 Memory Available: %.2f GB", float64(memoryStats.Available)/1e9)
		if !detectedComponents[memInfo] {
			components = append(components, memInfo)
			detectedComponents[memInfo] = true
			fmt.Println("✅ Detected:", memInfo) // 🔹 Debug Log
		}
	}

	// Detect available disk space
	diskUsage, err := disk.Usage("/")
	if err == nil {
		diskInfo := fmt.Sprintf("📀 Disk Space Available: %.2f GB", float64(diskUsage.Free)/1e9)
		if !detectedComponents[diskInfo] {
			components = append(components, diskInfo)
			detectedComponents[diskInfo] = true
			fmt.Println("✅ Detected:", diskInfo) // 🔹 Debug Log
		}
	}

	// Scan for system files
	entries, err := os.ReadDir(".")
	if err == nil {
		for _, entry := range entries {
			if strings.HasSuffix(entry.Name(), ".json") || strings.HasSuffix(entry.Name(), ".yaml") {
				configFile := fmt.Sprintf("⚙️ Config File: %s", entry.Name())
				if !detectedComponents[configFile] {
					components = append(components, configFile)
					detectedComponents[configFile] = true
					fmt.Println("✅ Detected:", configFile) // 🔹 Debug Log
				}
			}
		}
	}

	if len(components) == 0 {
		fmt.Println("✅ No new components detected.")
	} else {
		fmt.Println("📌 Newly Detected Components:", components)
	}

	return components
}

// Get network latency
func getNetworkLatency(server string) float64 {
	start := time.Now()
	conn, err := net.Dial("udp", server)
	if err != nil {
		fmt.Println("❌ Network latency error:", err)
		return 0.0
	}
	defer conn.Close()
	return time.Since(start).Seconds() * 1000 // Convert to ms
}

// Mock CPU temperature function
func getTemperature() float64 {
	return 65.0 // Placeholder temperature
}

// GenerateVector converts text into a hashed float32 vector
func GenerateVector(data string) []float32 {
	hash := sha256.Sum256([]byte(data))
	vector := make([]float32, 3)

	for i := 0; i < 3; i++ {
		vector[i] = float32(binary.BigEndian.Uint32(hash[i*8:(i+1)*8])) / 1e9
	}
	return vector
}

// RememberEvent saves recent events (short-term memory)
func rememberEvent(event string) {
	const maxEvents = 10
	if len(recentEvents) >= maxEvents {
		recentEvents = recentEvents[1:]
	}
	recentEvents = append(recentEvents, event)
}

// ReviewRecentEvents periodically reviews short-term memory
func reviewRecentEvents() {
	fmt.Println("🧠 Reviewing recent events in short-term memory...")
	for i, event := range recentEvents {
		fmt.Printf("📌 Event %d: %s\n", i+1, event)
	}
}
