package claim

import "charm.land/bubbletea/v2"

func Wait(updates chan Message) tea.Cmd {
	return func() tea.Msg {
		return <-updates
	}
}
