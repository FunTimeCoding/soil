package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"time"
)

func (s *Service) FlushDrained(drained map[string][]*event.Event) {
	now := time.Now()

	for root, events := range drained {
		s.report(s.flush(root, events, now))
	}
}
