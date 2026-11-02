package service

import "github.com/funtimecoding/soil/pkg/keepass"

func (s *Service) discard() {
	c, e := keepass.New(s.path, s.password)

	if e != nil {
		s.reportReload(e)

		return
	}

	s.client = c
}
