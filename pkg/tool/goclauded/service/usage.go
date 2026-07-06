package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/site/usage_result"

func (s *Service) Usage() *usage_result.Result {
	s.usageMutex.RLock()
	defer s.usageMutex.RUnlock()

	return s.usage
}
