package metric

func (m *Metric) CountSweptPlacements(
	source string,
	swept int64,
) {
	m.sweptPlacements.WithLabelValues(source).Add(float64(swept))
}
