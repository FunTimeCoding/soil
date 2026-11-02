package usage

import "github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"

func (s *Service) UtilizationCredential() *credential.Credential {
	if s.utilization == nil {
		return nil
	}

	return s.credential()
}
