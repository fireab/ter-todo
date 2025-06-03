package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type StatusType string

const (
	ToDo    StatusType = "todo"
	Done    StatusType = "done"
	Pending StatusType = "pending"
)

type FocusType int

const (
	TaskInputFocus FocusType = 1
	DesInputFocus  FocusType = 2
	TableFocus     FocusType = 3
)

type statusStates []StatusType

type taskInput struct {
	Id          uuid.UUID
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

	ti2 := textinput.New()
	ti2.Placeholder = "Description"
	ti2.Cursor.Blink = true
	ti2.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F0F8FF")).
		Bold(true).
		Blink(true).
		Underline(true)

	ti2.Prompt = "✏️  "
	ti2.CharLimit = 400
	ti2.Width = 100

	var columns = []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Title", Width: 30},
		{Title: "Description", Width: 50},
		{Title: "Status", Width: 10},
	}

	// initalize table
	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	// s.Cell = s.Cell.BorderStyle(lipgloss.RoundedBorder())
	t.SetStyles(s)

	return model{
		taskInput:       ti,
		descriptonInput: ti2,
		TaskList:        []taskInput{},
		err:             nil,
		table:           t,
		focusInput:      TaskInputFocus,
		stateStatus:     []StatusType{ToDo, Pending, Done},
	}
}

// Define the model
type model struct {
	taskInput       textinput.Model
	descriptonInput textinput.Model
	TaskList        []taskInput
	hight           int
	width           int
	err             error
	table           table.Model
	focusInput      FocusType
	stateStatus     statusStates
}

// Initialize the model
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) UpdateStates(index int, key string) StatusType {
	if key == tea.KeyRight.String() {
		if index < len(m.stateStatus)-1 {
			return m.stateStatus[index+1]
		} else {
			return m.stateStatus[0]
		}
	} else {
		if index > 0 {
			return m.stateStatus[index-1]
		} else {
			return m.stateStatus[len(m.stateStatus)-1]
		}
	}

}

func (m model) DeleteTask(index int) {
	if index < 0 || index > len(m.TaskList)-1 {
		return
	}
	m.TaskList = append(m.TaskList[:index], m.TaskList[index+1:]...)

}

func arrayToRow(m model, index int, task taskInput) {

}
