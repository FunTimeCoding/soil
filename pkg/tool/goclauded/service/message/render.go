package message

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/delivery_renderer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Service) Render(entries []queue.Entry) (string, error) {
	messages, e := s.messagesFor(entries)

	if e != nil {
		return "", e
	}

	return delivery_renderer.New(messages, time.Local).Context(entries), nil
}
