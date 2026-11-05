package service

import "github.com/funtimecoding/soil/pkg/iterm/screen"

func (s *Service) ReadScreen(identifier string) (*screen.Screen, error) {
	return guarded(
		s.guard,
		func() (*screen.Screen, error) {
			return s.client.ReadScreen(identifier)
		},
	)
}
