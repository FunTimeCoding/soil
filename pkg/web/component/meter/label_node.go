package meter

import "maragu.dev/gomponents"

func (m *Meter) LabelNode(n ...gomponents.Node) *Meter {
	m.label = append(m.label, n...)

	return m
}
