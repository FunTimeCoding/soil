package monitor

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/system"
)

func (m *Model) keyEvent(g tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch g.String() {
	case constant.KeyEscape:
		if m.modal != nil {
			m.modal = nil

			return m, nil
		}

		if m.table.Focused() {
			m.table.Blur()
		} else {
			m.table.Focus()
		}
	case constant.KeyQ, constant.KeyCtrlC:
		return m, tea.Quit
	case constant.KeyD:
		return m, viewDetail()
	case constant.KeyO:
		system.OpenBrowser(m.selectedItem().Link)
	case constant.KeyR:
		if !m.auto {
			return m, fetch.Command()
		}
	case constant.KeyM:
		if m.auto {
			m.auto = false
		} else {
			m.auto = true
		}

		return m, nil
	case constant.KeyEnter:
		return m, m.toggleClaim()
	case "t":
		return m, tea.Batch(addToast("Pressed t"))
	}

	return nil, nil
}
