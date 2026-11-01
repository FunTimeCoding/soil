package meter

func (m *Meter) Width(width string) *Meter {
	m.width = width

	return m
}
