package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"

func (s *Service) Decision(identifier uint) *decision.Decision {
	return s.store.Decision(identifier)
}
