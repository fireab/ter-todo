package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	var taskListStyle = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#1F1D2B")).
		Foreground(lipgloss.Color("#7FFFD4")).
		Bold(true).
		PaddingLeft(0).
		MarginLeft(10).
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Left).
		Width(m.width - 20)

	label := lipgloss.NewStyle().
		Bold(true).
		MarginLeft(10).
		Width(12).
		Render("Enter Task:")

	input := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Background(lipgloss.Color("#1F1D2B")).
		Padding(0, 1).
		Width(30).
		Render(m.textInput.View())

	view := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7FFFD4")).
		MarginLeft(15).
		Align(lipgloss.Center).
		Render("TASKS")
	view += "\n"
	view += taskListStyle.Render(fmt.Sprintf("%s\n\n", formattedString(m.Tasks)))

	inlineRow := lipgloss.JoinHorizontal(lipgloss.Top, label, input)

	view += "\n"
	view += inlineRow

	view += taskListStyle.Render(fmt.Sprintf("%s\n", m.TaskList))

	if m.err != nil {
		view += fmt.Sprintf("Error: %v\n", m.err)
	}

	return view
}
