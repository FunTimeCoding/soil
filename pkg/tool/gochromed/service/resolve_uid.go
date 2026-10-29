package service

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (s *Service) resolveUID(
	tabIdentifier string,
	uid string,
) (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if v, okay := s.snapshotCache[tabIdentifier][uid]; okay {
		return v, nil
	}

	return 0, not_found.Format("uid %s not found - take a snapshot first", uid)
}
