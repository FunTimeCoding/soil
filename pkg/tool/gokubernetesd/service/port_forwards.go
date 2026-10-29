package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"
	"maps"
	"slices"
)

func (s *Service) PortForwards() []*port_forward.Forward {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return slices.Collect(maps.Values(s.portForwards))
}
