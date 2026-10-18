package service

import "github.com/funtimecoding/soil/pkg/peer"

func (s *Service) Guard() *peer.Guard {
	return s.guard
}
