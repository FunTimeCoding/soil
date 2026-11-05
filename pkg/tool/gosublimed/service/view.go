package service

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (s *Service) View(identifier int) (*view.View, error) {
	return guarded(
		s.guard,
		func() (*view.View, error) {
			return s.client.View(identifier)
		},
	)
}
