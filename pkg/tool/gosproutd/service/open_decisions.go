package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"

func (s *Service) OpenDecisions(frameIdentifier uint) []*decision.Decision {
	return s.store.OpenDecisions(frameIdentifier)
}
