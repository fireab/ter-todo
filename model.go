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
		// TaskList:    GetTasks(),
		err:         nil,
		table:       t,
		focusInput:  TaskInputFocus,
		stateStatus: []StatusType{ToDo, Pending, Done},
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

// add task to the db and the json
func (m *model) AddTask(task Task) error {
	response := DB.Create(&task)
	if response.Error != nil {
		err := response.Error
		return err
	}

	m.TaskList = append(m.TaskList, taskInput{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Descripton,
		Status:      StatusType(task.Status),
	})
	return nil
}

// get all tasks from the database
func (m model) GetTasks() ([]taskInput, error) {
	var tasks []Task
	response := DB.Find(&tasks)
	if response.Error != nil {
		err := response.Error
		return nil, err
	}
	var intialTaskList = make([]taskInput, len(tasks))
	for i, task := range tasks {
		intialTaskList[i] = taskInput{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Descripton,
			Status:      StatusType(task.Status),
		}
	}
	return intialTaskList, nil
}

// delete task form the db and json
func (m *model) DeleteTaskFromDB(taskId uuid.UUID) error {
	response := DB.Delete(&Task{}, taskId)
	if response.Error != nil {
		err := response.Error
		return err
	}

	// Remove the task from the TaskList
	for i, task := range m.TaskList {
		if task.Id == taskId {
			m.TaskList = append(m.TaskList[:i], m.TaskList[i+1:]...)
			break
		}
	}
	return nil
}

// update task in the db and json
func (m *model) UpdateTaskInDB(id uuid.UUID, task Task) error {
	// Find the task by ID
	task.ID = id
	response := DB.Save(&task)
	if response.Error != nil {
		err := response.Error
		return err
	}

	// Update the task in the TaskList
	for i, t := range m.TaskList {
		if t.Id == task.ID {
			m.TaskList[i] = taskInput{
				Id:          task.ID,
				Title:       task.Title,
				Description: task.Descripton,
				Status:      StatusType(task.Status),
			}
			break
		}
	}
	return nil
}
