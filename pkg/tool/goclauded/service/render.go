package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/delivery"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Service) Render(entries []queue.Entry) (string, error) {
	messages, e := s.messagesFor(entries)

	if e != nil {
		return "", e
	}

	return delivery.New(messages, time.Local).Context(entries), nil
}
