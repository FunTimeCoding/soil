package service

import "github.com/funtimecoding/soil/pkg/peer"

func (s *Service) MemoryGuard() *peer.Guard {
	return s.memoryGuard
}
