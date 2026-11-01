package meter

func (m *Meter) Shown(text string) *Meter {
	m.shown = text

	return m
}
