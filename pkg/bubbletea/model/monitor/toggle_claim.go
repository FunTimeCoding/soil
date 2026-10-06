package monitor

import (
	"charm.land/bubbletea/v2"
	"slices"
)

func (m *Model) toggleClaim() tea.Cmd {
	i := m.selectedItem()

	if m.monitor == nil || i == nil {
		return nil
	}

	c := m.monitor
	owner := m.owner
	identifier := i.Identifier
	held := slices.Contains(m.claims[identifier], owner)

	return func() tea.Msg {
		var e error

		if held {
			e = c.Release(identifier, owner)
		} else {
			e = c.ClaimItem(identifier, owner)
		}

		if e != nil {
			return addToastMessage(e.Error())
		}

		return nil
	}
}
