package monitor

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
)

func (m *Model) Init() tea.Cmd {
	return tick.Command()
}
