package service

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"

func (s *Service) StorePortForward(f *port_forward.Forward) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.portForwards[f.Identifier] = f
}
