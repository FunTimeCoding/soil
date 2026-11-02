package service

import "github.com/funtimecoding/soil/pkg/keepass"

func (s *Service) refresh() {
	if !s.client.Changed() {
		return
	}

	c, e := keepass.New(s.path, s.password)

	if e != nil {
		s.reportReload(e)

		return
	}

	s.logger.Structured("database changed on disk - reloaded")
	s.client = c
	s.reported = false
}
