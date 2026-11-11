package message

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/send_result"
	"strings"
)

func (s *Service) Send(
	name string,
	to []string,
	body string,
	requested constant.Delivery,
) (*send_result.Result, error) {
	requested, e := delivery(requested, constant.DeliveryImmediate)

	if e != nil {
		return nil, e
	}

	holders, f := s.recipients(name, to)

	if f != nil {
		return nil, f
	}

	d := requested

	if len(to) == 0 {
		d = constant.DeliveryQueued
	}

	m, g := s.store.SendMessage(name, strings.Join(to, ", "), body)

	if g != nil {
		return nil, g
	}

	formatted := fmt.Sprintf("%s: %s", name, body)
	result := send_result.New(requested, d)

	for _, h := range holders {
		entry := queue.New(
			h.Identifier,
			h.CallsignValue(),
			constant.QueueMessage,
			formatted,
		)
		entry.MessageIdentifier = new(m.Identifier)
		delivered, i := s.pushDelivery(entry, d)

		if i != nil {
			return nil, i
		}

		if d == constant.DeliveryImmediate && !delivered {
			result.Held = append(result.Held, h.CallsignValue())
		}
	}

	return result, nil
}
