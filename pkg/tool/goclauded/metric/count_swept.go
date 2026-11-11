package metric

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (m *Metric) CountSwept(
	collection string,
	swept int,
) {
	m.reconciled.WithLabelValues(
		collection,
		constant.OutcomeSwept,
	).Add(float64(swept))
}
