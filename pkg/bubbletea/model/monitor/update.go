package monitor

import (
	"charm.land/bubbletea/v2"
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/receive"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/toast"
	"time"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.WindowSizeMsg:
		m.resizeEvent(g)
	case tea.KeyMsg:
		if o, c := m.keyEvent(g); o != nil {
			return o, c
		}
	case tick.Message:
		return m.tickEvent(g)
	case fetch.Message:
		m.fetchEvent(g)
	case receive.Message:
		m.receiveEvent(g)
	case addToastMessage:
		n := toast.New(m.nextToast, string(g))
		m.toast = append(m.toast, n)
		m.nextToast++
		m.updateTableHeight(false, false)
		m.table.UpdateViewport()

		return m, removeToast(n.Identifier, 10*time.Second)
	case removeToastMessage:
		var updated []*toast.Toast

		for _, n := range m.toast {
			if n.Identifier != int(g) {
				updated = append(updated, n)
			}
		}

		before := len(m.toast)
		m.toast = updated
		after := len(m.toast)
		m.updateTableHeight(before > 0 && after == 0, false)
		m.table.UpdateViewport()

		return m, nil
	case viewDetailMessage:
		i := m.selectedItem()
		m.modal = &Modal{
			content: fmt.Sprintf("Details: %+v", i),
		}
	}

	t, result := m.table.Update(s)
	m.table = &t

	return m, result
}
