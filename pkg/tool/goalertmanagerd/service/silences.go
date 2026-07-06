package service

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"

func (s *Service) Silences(
	instance string,
	expired bool,
) ([]*silence.Silence, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Silences(expired)
}
