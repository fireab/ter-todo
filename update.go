package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func getIndex(value StatusType, arr statusStates) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1 // not found
}

var val = 0

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if val == 0 {
		m.TaskList, _ = m.GetTasks()

		var rows = make([]table.Row, len(m.TaskList))
		for i, task := range m.TaskList {
			rows[i] = table.Row{
				strconv.Itoa(i + 1),
				task.Title,
				task.Description,
				string(task.Status),
			}
		}
		m.table.SetRows(rows)
		val = 1
	}
	if m.focusInput == TaskInputFocus {
		m.taskInput, cmd = m.taskInput.Update(msg)
	} else if m.focusInput == DesInputFocus {
		m.descriptonInput, cmd = m.descriptonInput.Update(msg)
	} else {
		// fmt.Println(">>", m.TaskList)
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
					m.AddTask(Task{
						ID:         uuid.New(),
						Descripton: desc,
						Title:      val,
						Status:     string(ToDo),
					})

					var rows = make([]table.Row, len(m.TaskList))
					for i, task := range m.TaskList {
						rows[i] = table.Row{
							strconv.Itoa(i + 1),
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

			return m, nil
		case tea.KeyLeft.String(), tea.KeyRight.String():
			if m.focusInput == TableFocus {
				rowStatus := m.table.SelectedRow()[3]
				rowIdInt, err := strconv.Atoi(m.table.SelectedRow()[0])
				if err != nil {
					fmt.Println("Error converting rowId to integer:", err)
					return m, nil
				}

				index := getIndex(StatusType(rowStatus), m.stateStatus)
				s := m.UpdateStates(index, msg.String())

				m.TaskList[rowIdInt-1].Status = s
				// let render the array to the table
				var rows = make([]table.Row, len(m.TaskList))
				for i, task := range m.TaskList {
					rows[i] = table.Row{
						strconv.Itoa(i + 1),
						task.Title,
						task.Description,
						string(task.Status),
					}
				}
				m.table.SetRows(rows)

			}
			return m, nil
		// implement delete
		case tea.KeyDelete.String():
			if m.focusInput == TableFocus {
				if len(m.TaskList) == 0 {
					return m, nil
				}
				rowId := m.table.SelectedRow()[0]
				// remove index rowId-1 from TaskList
				rowIdInt, err := strconv.Atoi(rowId)
				if err != nil {
					fmt.Println("Error converting rowId to integer:", err)
					return m, nil
				}

				if rowIdInt > 0 && rowIdInt <= len(m.TaskList) {
					m.TaskList = append(m.TaskList[:rowIdInt-1], m.TaskList[rowIdInt:]...)
					// let render the array to the table
					var rows = make([]table.Row, len(m.TaskList))
					for i, task := range m.TaskList {
						rows[i] = table.Row{
							strconv.Itoa(i + 1),
							task.Title,
							task.Description,
							string(task.Status),
						}
					}
					m.focusInput = TableFocus
					m.table.MoveUp(1)
					m.table.SetRows(rows)
				} else {
					fmt.Println("Invalid row ID:", rowIdInt)
					return m, nil
				}

			}

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
