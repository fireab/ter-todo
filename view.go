package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	label := lipgloss.NewStyle().
		Bold(true).
		MarginLeft(10).
		Width(12).
		Height(1).
		Render("Enter Task:")

	input := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Background(lipgloss.Color("#1F1D2B")).
		Padding(0, 1).
		Width(30).
		Italic(true).
		Render(m.taskInput.View())

	input2 := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Background(lipgloss.Color("#1F1D2B")).
		Padding(0, 1).
		Width(50).
		Italic(true).
		Render(m.descriptonInput.View())

	inlineRow := lipgloss.JoinHorizontal(lipgloss.Top, label, input)
	inlineRow2 := lipgloss.JoinHorizontal(lipgloss.Top, label, input2)
	var baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		MarginLeft(10)

	var view = baseStyle.Render(m.table.View()) + "\n"
	view += "\n"
	view += inlineRow
	view += "\n"
	view += inlineRow2
	view += "\n"

	if m.err != nil {
		view += fmt.Sprintf("Error: %v\n", m.err)
	}
	// create table for TaskLIsts

	return view
}
