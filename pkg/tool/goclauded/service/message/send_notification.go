package message

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Service) SendNotification(
	callsign string,
	source string,
	body string,
	requested constant.Delivery,
) (bool, error) {
	d, e := delivery(requested, constant.DeliveryQueued)

	if e != nil {
		return false, e
	}

	if f := entryRefusal(constant.QueueNotification, body); f != nil {
		return false, f
	}

	holder, e := s.store.CallsignSession(callsign)

	if e != nil {
		return false, e
	}

	if holder == nil {
		return false, not_found.New(constant.Callsign, callsign)
	}

	if e := s.store.SendNotification(
		holder.Identifier,
		callsign,
		source,
		body,
	); e != nil {
		return false, e
	}

	return s.pushDelivery(
		queue.New(
			holder.Identifier,
			callsign,
			constant.QueueNotification,
			fmt.Sprintf("%s: %s", source, body),
		),
		d,
	)
}
