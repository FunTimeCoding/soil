package meter

func (m *Meter) Class(class string) *Meter {
	m.class = class

	return m
}
