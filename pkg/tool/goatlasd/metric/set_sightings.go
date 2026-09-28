package metric

func (m *Metric) SetSightings(
	source string,
	count int,
) {
	m.sightings.WithLabelValues(source).Set(float64(count))
}
