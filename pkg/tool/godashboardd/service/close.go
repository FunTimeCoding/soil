package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.prometheusGuard.Close()
			s.prometheus.Close()

			if s.argocd != nil {
				s.argocdGuard.Close()
				s.argocd.Close()
			}

			if s.usage != nil {
				s.usageGuard.Close()
				s.usage.Close()
			}

			s.authorization.Close()
			relational.CloseMapper(s.database)
		},
	)
}
