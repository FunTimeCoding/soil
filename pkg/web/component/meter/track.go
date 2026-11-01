package meter

func (m *Meter) Track(color string) *Meter {
	m.track = color

	return m
}
