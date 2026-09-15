package model

import "charm.land/bubbletea/v2"

func (m *Model) View() tea.View {
	v := tea.NewView("")
	v.KeyboardEnhancements.ReportEventTypes = true

	return v
}
