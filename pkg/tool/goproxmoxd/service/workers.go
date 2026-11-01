package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Workers() []face.Worker {
	if s.netboxGuard == nil {
		return nil
	}

	return []face.Worker{s.netboxGuard}
}
