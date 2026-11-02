package service

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"

func (s *Service) FlushDrained(drained map[string][]*event.Event) {
	for root, events := range drained {
		s.report(s.flush(root, events))
	}
}
