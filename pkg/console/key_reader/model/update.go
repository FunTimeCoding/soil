package model

import (
	"charm.land/bubbletea/v2"
	bubbleteaConstant "github.com/funtimecoding/soil/pkg/bubbletea/constant"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.KeyPressMsg:
		switch g.String() {
		case bubbleteaConstant.KeyEscape, bubbleteaConstant.KeyCtrlC:
			return m, tea.Quit
		}

		if !g.IsRepeat {
			m.reader.Press(g.Code)
		}
	case tea.KeyReleaseMsg:
		m.reader.Release(g.Code)
	}

	return m, nil
}
