package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"

func (s *Service) Decisions(session string) []*decision.Decision {
	return s.store.Decisions(session)
}
