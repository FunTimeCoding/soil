package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Service) pushDelivery(
	e *queue.Entry,
	d constant.Delivery,
) (bool, error) {
	e.Delivery = d

	switch d {
	case constant.DeliveryImmediate:
		return s.pushImmediate(e)
	case constant.DeliveryCruise:
		return s.pushCruise(e)
	default:
		return false, s.pushEntry(e)
	}
}
