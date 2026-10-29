package service

import "github.com/funtimecoding/soil/pkg/iterm/screen"

func (s *Service) ReadScreen(identifier string) (*screen.Screen, error) {
	return s.source.ReadScreen(identifier)
}
