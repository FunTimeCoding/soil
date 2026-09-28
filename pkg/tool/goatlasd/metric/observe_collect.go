package metric

import "time"

func (m *Metric) ObserveCollect(
	source string,
	took time.Duration,
) {
	m.seconds.WithLabelValues(source).Observe(took.Seconds())
}
