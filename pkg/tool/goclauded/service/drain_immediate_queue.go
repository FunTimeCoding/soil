package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/delivery"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Service) DrainImmediateQueue(
	sessionIdentifier string,
	callsign string,
) ([]queue.Entry, error) {
	result, e := s.store.DrainImmediateQueue(sessionIdentifier, callsign)

	if e != nil {
		return nil, e
	}

	messages, f := s.messagesFor(result)

	if f != nil {
		return nil, f
	}

	d := delivery.New(messages, time.Local)

	for i, v := range result {
		if v.Kind == constant.QueueMessage {
			result[i].Body = d.Cut(v)
		}
	}

	return result, nil
}
