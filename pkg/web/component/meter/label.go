package meter

import "maragu.dev/gomponents"

func (m *Meter) Label(text string) *Meter {
	return m.LabelNode(gomponents.Text(text))
}
