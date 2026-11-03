package metric

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (m *Metric) CountReconcile(
	collection string,
	pushed int,
	failed int,
) {
	m.reconciled.WithLabelValues(
		collection,
		constant.OutcomePushed,
	).Add(float64(pushed))
	m.reconciled.WithLabelValues(
		collection,
		constant.OutcomeFailed,
	).Add(float64(failed))
}
