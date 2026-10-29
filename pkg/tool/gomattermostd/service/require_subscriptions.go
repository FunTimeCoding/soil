package service

import "github.com/funtimecoding/soil/pkg/errors/not_configured"

func (s *Service) requireSubscriptions() error {
	if s.watcher == nil {
		return not_configured.Format("thread subscriptions need goclauded")
	}

	return nil
}
