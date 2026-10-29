package service

import "github.com/funtimecoding/soil/pkg/errors/not_configured"

func (s *Service) RunMonitor() error {
	if s.monitor == nil {
		return not_configured.Format("topic monitoring is not configured")
	}

	return s.monitor.RunNow()
}
