package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.focusInput == TaskInputFocus {
		m.taskInput, cmd = m.taskInput.Update(msg)
	} else if m.focusInput == DesInputFocus {

		m.descriptonInput, cmd = m.descriptonInput.Update(msg)
	} else {

		m.table, cmd = m.table.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			if m.focusInput == TaskInputFocus || m.focusInput == DesInputFocus {
				val := strings.TrimSpace(m.taskInput.Value())
				desc := strings.TrimSpace(m.descriptonInput.Value())
				if val != "" {
					m.TaskList = append(m.TaskList, taskInput{
						Id:          len(m.TaskList) + 1,
						Title:       val,
						Description: desc,
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
				m.taskInput.Reset()
				m.taskInput.Focus()
				m.descriptonInput.Blur()
				m.table.Blur()
				m.focusInput = TaskInputFocus
				m.descriptonInput.Reset()
			}
			if m.focusInput == TableFocus {
				// fmt.Println("Table clicked", m.table.SelectedRow()[1])
				// fmt.Println("Table Focuse", m.table.Focused())
				return m, tea.Batch(
					tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
				)
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
			if m.focusInput == TaskInputFocus {
				m.focusInput = DesInputFocus
				m.table.Blur()
				m.taskInput.Blur()
				m.descriptonInput.Focus()
				return m, nil

			} else if m.focusInput == DesInputFocus {
				m.focusInput = TableFocus
				m.taskInput.Blur()
				m.table.Blur()
				m.table.Focus()
				return m, nil

			} else if m.focusInput == TableFocus {
				m.focusInput = TaskInputFocus
				m.taskInput.Focus()
				m.table.Blur()
				m.descriptonInput.Blur()
				return m, nil

			} else {
				return m, nil
			}

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
