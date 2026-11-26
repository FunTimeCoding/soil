package service

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (s *Service) resolveUID(
	tab string,
	uid string,
) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if v, okay := s.references[tab][uid]; okay {
		return v, nil
	}

	return "", not_found.Format("uid %s not found - take a snapshot first", uid)
}
