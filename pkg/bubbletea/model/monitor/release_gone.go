package monitor

import (
	"charm.land/bubbletea/v2"
	"slices"
)

func (m *Model) releaseGone() tea.Cmd {
	if m.monitor == nil {
		return nil
	}

	present := map[string]bool{}

	for _, i := range m.items {
		present[i.Identifier] = true
	}

	var gone []string

	for identifier, owners := range m.claims {
		if !present[identifier] && slices.Contains(owners, m.owner) {
			gone = append(gone, identifier)
		}
	}

	if len(gone) == 0 {
		return nil
	}

	c := m.monitor
	owner := m.owner

	return func() tea.Msg {
		for _, identifier := range gone {
			if e := c.Release(identifier, owner); e != nil {
				return addToastMessage(e.Error())
			}
		}

		return nil
	}
}
