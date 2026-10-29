package service

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"

func (s *Service) forwardEnded(
	f *port_forward.Forward,
	e error,
) {
	s.mutex.Lock()

	if s.portForwards[f.Identifier] == f {
		delete(s.portForwards, f.Identifier)
	}

	s.mutex.Unlock()

	if e != nil {
		s.reporter.CaptureException(e)
	}
}
