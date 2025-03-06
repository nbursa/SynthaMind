package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Task struct {
	ID      int
	Name    string
	Vector  []byte
}

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to SQLite database
	var errDb error
	db, errDb = sql.Open("sqlite3", "./knowledge.db")
	if errDb != nil {
		log.Fatal("Failed to connect to database: ", errDb)
	}

	// Create the table for tasks
	_, errDb = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			task_name TEXT,
			vector BLOB
		);
	`)
	if errDb != nil {
		log.Fatal("Error creating table: ", errDb)
	}
}

func main() {
	// Start task manager
	fmt.Println("🚀 EvolvAI Task Manager Starting...")
	go startTaskManager()

	// Simulate incoming knowledge tasks
	for i := 0; i < 3; i++ {
		taskName := fmt.Sprintf("Sample task %d", i+1)
		insertTask(taskName, []byte{0x01, 0x02, 0x03, 0x04}) // Example vector data
		time.Sleep(2 * time.Second)
	}

	// Fetch all tasks from the database
	fetchTasks()
}

// Function to start task manager (for now just a print statement)
func startTaskManager() {
	fmt.Println("🧠 Task Manager Running...")
}

// Function to insert a task into the database
func insertTask(taskName string, vector []byte) {
	statement, err := db.Prepare("INSERT INTO tasks (task_name, vector) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
	}
	defer statement.Close()

	_, err = statement.Exec(taskName, vector)
	if err != nil {
		log.Fatal("Error inserting task: ", err)
	}

	fmt.Printf("🟢 Task '%s' inserted into database\n", taskName)
}

// Function to fetch all tasks from the database
func fetchTasks() {
	rows, err := db.Query("SELECT id, task_name FROM tasks")
	if err != nil {
		log.Fatal("Error fetching tasks: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name)
		if err != nil {
			log.Fatal("Error scanning task: ", err)
		}
		fmt.Printf("ID: %d, Task: %s\n", task.ID, task.Name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal("Error during row iteration: ", err)
	}
}
