package constant

import "charm.land/lipgloss/v2"

const LogFile = "tea.txt"

var (
	Default = lipgloss.NewStyle()

	Table = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(
		lipgloss.Color("240"),
	)

	Modal = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Width(50).Height(
		10,
	).Align(
		lipgloss.Center,
	)
)
