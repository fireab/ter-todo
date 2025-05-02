package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatusType string

const (
	ToDo    StatusType = "todo"
	Done    StatusType = "done"
	Pending StatusType = "pending"
)

type taskInput struct {
	Title       string
	Description string
	Status      StatusType
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Task"
	ti.Cursor.Blink = true
	ti.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F0F8FF")).
		Bold(true).
		Blink(true).
		Underline(true)

	ti.Prompt = "✏️  "
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		Tasks:     []string{},
		Message:   "",
		TaskList:  []taskInput{},
		err:       nil,
	}
}

// Define the model
type model struct {
	Tasks     []string
	Message   string
	textInput textinput.Model
	TaskList  []taskInput
	hight     int
	width     int
	err       error
}

// Initialize the model
func (m model) Init() tea.Cmd {

	return textinput.Blink
}
