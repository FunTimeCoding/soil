package meter

func (m *Meter) Color(color string) *Meter {
	m.color = color

	return m
}
