package service

import "github.com/funtimecoding/soil/pkg/errors/not_configured"

func (s *Service) requireWatches() error {
	if s.store == nil {
		return not_configured.Format("merge request watches need goclauded")
	}

	return nil
}
