package monitor

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
)

func (m *Model) Init() tea.Cmd {
	result := []tea.Cmd{tick.Command()}

	if m.monitor != nil {
		go claim.Watch(m.monitor, m.updates)
		result = append(result, claim.Wait(m.updates))
	}

	if m.notice != "" {
		result = append(result, addToast(m.notice))
	}

	return tea.Batch(result...)
}
