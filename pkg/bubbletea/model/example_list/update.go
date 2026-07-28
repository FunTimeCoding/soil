package example_list

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.KeyMsg:
		switch g.String() {
		case constant.KeyCtrlC, constant.KeyQ:
			return m, tea.Quit
		case constant.KeyUp, constant.KeyK:
			if m.cursor > 0 {
				m.cursor--
			}
		case constant.KeyDown, constant.KeyJ:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case constant.KeyEnter, constant.KeySpace:
			if _, okay := m.selected[m.cursor]; okay {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}
