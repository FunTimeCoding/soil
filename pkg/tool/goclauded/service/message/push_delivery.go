package message

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Service) pushDelivery(
	e *queue.Entry,
	d constant.Delivery,
) (bool, error) {
	e.Delivery = d

	if d == constant.DeliveryImmediate {
		return s.pushImmediate(e)
	}

	return false, s.pushEntry(e)
}
