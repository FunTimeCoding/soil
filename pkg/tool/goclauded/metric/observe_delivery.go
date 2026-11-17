package metric

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (m *Metric) ObserveDelivery(characters int) bool {
	m.delivered.Observe(float64(characters))

	if characters <= constant.DeliveryCharacterLimit {
		return false
	}

	m.overLimit.Inc()

	return true
}
