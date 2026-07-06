package example_list

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/key"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.KeyMsg:
		switch g.String() {
		case key.CtrlC, key.Q:
			return m, tea.Quit
		case key.Up, key.K:
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Down, key.J:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case key.Enter, key.Space:
			if _, okay := m.selected[m.cursor]; okay {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}
