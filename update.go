package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Update the text input first
	m.textInput, cmd = m.textInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case tea.KeyEnter.String():
			if strings.TrimSpace(m.textInput.Value()) != "" {
				m.Tasks = append(m.Tasks, m.textInput.Value())

				m.TaskList = append(m.TaskList, taskInput{
					Title:       m.textInput.Value(),
					Description: "description",
					Status:      ToDo,
				})

				m.textInput.Reset()
				m.textInput.Focus()
			}

		case tea.KeyCtrlC.String(), tea.KeyEsc.String():
			return m, tea.Quit

		default:
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.hight = msg.Height

	case error:
		m.err = msg
		return m, nil
	}

	return m, cmd
}
