package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"

func (s *Service) UtilizationCredential() *utilization.Credential {
	if !s.UtilizationFallback() {
		return nil
	}

	return utilization.ReadCredential()
}
