package monitor

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
)

func (m *Model) claimEvent(g claim.Message) tea.Cmd {
	if g.Error != nil {
		m.claimError = g.Error

		return claim.Wait(m.updates)
	}

	m.claimError = nil
	m.claims = map[string][]string{}

	for _, c := range g.Claims {
		m.claims[c.Item] = append(m.claims[c.Item], c.Owner)
	}

	m.applyClaims()

	return claim.Wait(m.updates)
}
