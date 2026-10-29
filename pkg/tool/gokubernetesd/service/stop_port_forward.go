package service

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (s *Service) StopPortForward(identifier string) error {
	s.mutex.Lock()
	f, okay := s.portForwards[identifier]
	delete(s.portForwards, identifier)
	s.mutex.Unlock()

	if !okay {
		return not_found.New("port forward", identifier)
	}

	f.Stop()

	return nil
}
