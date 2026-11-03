package metric

func (m *Metric) SetIndexed(
	collection string,
	documents int,
) {
	m.indexed.WithLabelValues(collection).Set(float64(documents))
}
