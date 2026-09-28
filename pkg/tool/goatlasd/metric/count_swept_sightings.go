package metric

func (m *Metric) CountSweptSightings(
	source string,
	swept int64,
) {
	m.sweptSightings.WithLabelValues(source).Add(float64(swept))
}
