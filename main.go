package main

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"time"

	"evolvai/taskmanager"
	"evolvai/utils"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func insertData(db *sql.DB) {
	tasks := []utils.Task{
		{Data: "Sample task 1", Vector: []float32{0.1, 0.2, 0.3}},
		{Data: "Sample task 2", Vector: []float32{0.4, 0.5, 0.6}},
		{Data: "Sample task 3", Vector: []float32{0.7, 0.8, 0.9}},
	}

	for _, task := range tasks {
		// Check if task already exists to prevent duplicates
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE task_name = ?)", task.Data).Scan(&exists)
		if err != nil {
			log.Fatal(err)
		}
		if exists {
			fmt.Printf("⚠️ Task '%s' already exists, skipping insertion.\n", task.Data)
			continue
		}

		// Convert float32 slice to binary BLOB
		vectorBytes := make([]byte, len(task.Vector)*4)
		for i, v := range task.Vector {
			binary.LittleEndian.PutUint32(vectorBytes[i*4:], math.Float32bits(v))
		}

		// Insert task with binary vector
		_, err = db.Exec("INSERT INTO tasks (task_name, vector) VALUES (?, ?)", task.Data, vectorBytes)
		if err != nil {
			log.Fatal(err)
		}
		utils.LogInfo(fmt.Sprintf("🟢 Task '%s' inserted into database", task.Data))
	}
}

func fetchTasks(db *sql.DB) {
	rows, err := db.Query("SELECT id, task_name, vector FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var taskID int
		var taskName string
		var vectorData []byte

		err := rows.Scan(&taskID, &taskName, &vectorData)
		if err != nil {
			log.Fatal(err)
		}

		// Convert binary BLOB back to []float32
		vector := make([]float32, len(vectorData)/4)
		for i := range vector {
			vector[i] = math.Float32frombits(binary.LittleEndian.Uint32(vectorData[i*4:]))
		}

		fmt.Printf("📝 Task ID: %d, Task Name: %s, Vector: %v\n", taskID, taskName, vector)
	}
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database setup (SQLite example)
	database, err := sql.Open("sqlite3", "./knowledge.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Ensure the 'tasks' table exists
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_name TEXT,
		vector BLOB
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert example data
	insertData(database)

	// Fetch tasks from the database to verify
	fetchTasks(database)

	// Start task manager
	fmt.Println("🚀 EvolvAI Task Manager Starting...")
	go taskmanager.StartTaskManager()

	// Simulate incoming knowledge tasks
	for {
		taskmanager.AddTask("New knowledge detected")
		time.Sleep(5 * time.Second)
	}
}