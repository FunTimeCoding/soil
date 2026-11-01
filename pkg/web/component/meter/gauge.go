package meter

func (m *Meter) Gauge() *Meter {
	m.gauge = true

	return m
}
