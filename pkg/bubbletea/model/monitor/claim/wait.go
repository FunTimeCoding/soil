package claim

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
)

func Wait(updates chan message.Message) tea.Cmd {
	return func() tea.Msg {
		return <-updates
	}
}
