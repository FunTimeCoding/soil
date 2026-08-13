package style

import (
	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
)

func Table(m *table.Model) {
	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(
		lipgloss.Color("240"),
	).BorderBottom(true).Bold(false)
	s.Selected = s.Selected.Foreground(lipgloss.Color("229")).Background(
		lipgloss.Color("57"),
	).Bold(false)
	m.SetStyles(s)
}
