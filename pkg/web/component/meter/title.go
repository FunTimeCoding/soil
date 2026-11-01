package meter

func (m *Meter) Title(text string) *Meter {
	m.title = text

	return m
}
