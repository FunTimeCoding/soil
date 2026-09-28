package metric

func (m *Metric) CountFailure(
	source string,
	failed bool,
) {
	c := m.failures.WithLabelValues(source)

	if failed {
		c.Inc()
	}
}
