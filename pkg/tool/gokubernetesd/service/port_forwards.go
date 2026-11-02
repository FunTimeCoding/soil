package service

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"

func (s *Service) PortForwards() []*port_forward.Forward {
	return s.forwards.All()
}
