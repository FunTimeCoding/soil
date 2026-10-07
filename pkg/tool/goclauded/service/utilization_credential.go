package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
)

func (s *Service) UtilizationCredential() *credential.Credential {
	if !s.UtilizationFallback() {
		return nil
	}

	return utilization.ReadCredential()
}
