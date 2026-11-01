package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"strings"
)

func (s *Service) Send(
	name string,
	to []string,
	body string,
	d constant.Delivery,
) ([]string, error) {
	holders, e := s.recipients(name, to)

	if e != nil {
		return nil, e
	}

	if len(to) == 0 {
		d = constant.DeliveryQueued
	}

	m, f := s.store.SendMessage(name, strings.Join(to, ", "), body)

	if f != nil {
		return nil, f
	}

	formatted := fmt.Sprintf("%s: %s", name, body)
	var held []string

	for _, h := range holders {
		entry := queue.New(
			h.Identifier,
			h.CallsignValue(),
			constant.QueueMessage,
			formatted,
		)
		entry.MessageIdentifier = new(m.Identifier)
		delivered, g := s.pushDelivery(entry, d)

		if g != nil {
			return nil, g
		}

		if d != constant.DeliveryQueued && !delivered {
			held = append(held, h.CallsignValue())
		}
	}

	return held, nil
}
