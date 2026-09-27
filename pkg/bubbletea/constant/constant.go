package constant

import "charm.land/lipgloss/v2"

const (
	LogFile = "tea.txt"

	ItemIdentifierColumn = "Identifier"
	ItemScoreColumn      = "Score"
	ItemSeverityColumn   = "Severity"
	ItemDetailColumn     = "Detail"
	ItemUserColumn       = "User"

	KeyEnter  = "enter"
	KeyEscape = "esc"
	KeySpace  = "space"
	KeyUp     = "up"
	KeyDown   = "down"
	KeyD      = "d"
	KeyJ      = "j"
	KeyK      = "k"
	KeyL      = "l"
	KeyM      = "m"
	KeyO      = "o"
	KeyP      = "p"
	KeyQ      = "q"
	KeyR      = "r"
	KeyCtrlC  = "ctrl+c"
)

var (
	Default = lipgloss.NewStyle()
	Table   = lipgloss.NewStyle().BorderStyle(
		lipgloss.NormalBorder(),
	).BorderForeground(lipgloss.Color("240"))
	Modal = lipgloss.NewStyle().BorderStyle(
		lipgloss.NormalBorder(),
	).Width(50).Height(10).Align(
		lipgloss.Center,
	)
)
