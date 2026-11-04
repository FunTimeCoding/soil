package metric

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (m *Metric) ObserveDelivery(characters int) bool {
	m.delivered.Observe(float64(characters))

	if characters <= constant.DeliveryBudget {
		return false
	}

	m.overBudget.Inc()

	return true
}
