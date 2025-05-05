package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.focusInput {
		m.textInput, cmd = m.textInput.Update(msg)
	} else {
		m.table, cmd = m.table.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case tea.KeyEnter.String():
			if m.focusInput {

				val := strings.TrimSpace(m.textInput.Value())
				if val != "" {
					m.TaskList = append(m.TaskList, taskInput{
						Id:          len(m.TaskList) + 1,
						Title:       val,
						Description: "description",
						Status:      ToDo,
					})

					// ✅ Update table rows here
					var rows = make([]table.Row, len(m.TaskList))
					for i, task := range m.TaskList {
						rows[i] = table.Row{
							strconv.Itoa(task.Id),
							task.Title,
							task.Description,
							string(task.Status),
						}
					}
					m.table.SetRows(rows)
				}
				m.textInput.Reset()
			}
			return m, nil

		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
			return m, nil
		case "tab":
			// Toggle focus
			m.focusInput = !m.focusInput
			if m.focusInput {
				m.textInput.Focus()
				m.table.Blur()
			} else {
				m.textInput.Blur()
				m.table.Focus()
			}
			return m, nil

		case tea.KeyCtrlC.String():
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
