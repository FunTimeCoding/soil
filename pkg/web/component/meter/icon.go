package meter

import "maragu.dev/gomponents"

func (m *Meter) Icon(n gomponents.Node) *Meter {
	m.icon = n

	return m
}
