package monitor

import (
	"charm.land/bubbletea/v2"
	"context"
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
		x := context.Background()
		var e error

		if held {
			e = c.DeleteClaim(x, identifier, owner)
		} else {
			_, e = c.PutClaim(x, identifier, owner)
		}

		if e != nil {
			return addToastMessage(e.Error())
		}

		return nil
	}
}
