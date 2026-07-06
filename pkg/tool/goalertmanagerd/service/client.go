package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
)

func (s *Service) Client(instance string) (*alertmanager.Client, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, okay := s.clients[instance]; okay {
		return c, nil
	}

	i, okay := s.Instance(instance)

	if !okay {
		return nil, fmt.Errorf("unknown instance: %s", instance)
	}

	p := prometheus.New(
		i.Prometheus.Host,
		i.Prometheus.Port,
		i.Prometheus.Secure,
		i.Prometheus.User,
		i.Prometheus.Password,
		"",
	)
	c := alertmanager.New(
		i.Alertmanager.Host,
		i.Alertmanager.Secure,
		i.Alertmanager.User,
		i.Alertmanager.Password,
		p,
	)
	s.clients[instance] = c

	return c, nil
}
