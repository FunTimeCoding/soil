package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func (s *Service) Client(instance string) (*prometheus.Client, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, okay := s.clients[instance]; okay {
		return c, nil
	}

	i, okay := s.Instance(instance)

	if !okay {
		return nil, fmt.Errorf("unknown instance: %s", instance)
	}

	c := prometheus.New(
		i.Host,
		i.Port,
		i.Secure,
		i.User,
		i.Password,
		"",
	)
	s.clients[instance] = c

	return c, nil
}
