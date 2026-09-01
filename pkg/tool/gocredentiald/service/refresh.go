package service

import "github.com/funtimecoding/soil/pkg/keepass"

func (s *Service) refresh() {
	if !s.client.Changed() {
		return
	}

	s.logger.Structured("database changed on disk - reloading")
	s.client = keepass.New(s.path, s.password)
}
