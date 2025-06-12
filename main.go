package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// Import the pure Go SQLite driver
	_ "modernc.org/sqlite"
)

type Task struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Title      string    `gorm:"not null"`
	Descripton string    `gorm:"default:''"`
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func initDB() *gorm.DB {
	// Check if database file exists, create if not
	if _, err := os.Stat("task.db"); os.IsNotExist(err) {
		file, err := os.Create("task.db")
		if err != nil {
			panic("failed to create task.db file")
		}
		file.Close()
	}

	// Initialize DB connection using the pure Go SQLite driver
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic("failed to connect to database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Task{})
	if err != nil {
		fmt.Println("Error migrating database:", err)
		panic("failed to migrate database")
	}

	fmt.Println("Database connected and migrated!")
	return db
}

var DB *gorm.DB

func main() {
	DB = initDB()
	// create a task

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	// Start the program
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
