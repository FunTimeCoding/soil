package monitor

import (
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
)

func (m *Model) View() tea.View {
	var s string

	if m.width == 0 || m.height == 0 {
		s = "wait for screen size"
	} else if m.modal == nil {
		s = lipgloss.JoinVertical(
			lipgloss.Left,
			renderToasts(m.toast),
			constant.Default.Render(m.topBar),
			constant.Table.Render(m.table.View()),
			constant.Default.Render(m.bottomBar),
		)
	} else {
		s = lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			constant.Modal.Render(m.modal.content),
		)
	}

	v := tea.NewView(s)
	v.AltScreen = true

	return v
}
