package example_table

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.KeyMsg:
		switch g.String() {
		case constant.KeyEscape:
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case constant.KeyQ, constant.KeyCtrlC:
			return m, tea.Quit
		case constant.KeyEnter:
			return m, tea.Batch(
				tea.Printf(
					"Let's go to %s!",
					m.table.SelectedRow()[1],
				),
			)
		}
	}

	table, result := m.table.Update(s)
	m.table = &table

	return m, result
}
