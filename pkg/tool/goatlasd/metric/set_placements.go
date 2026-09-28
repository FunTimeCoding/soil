package metric

func (m *Metric) SetPlacements(
	source string,
	count int,
) {
	m.placements.WithLabelValues(source).Set(float64(count))
}
