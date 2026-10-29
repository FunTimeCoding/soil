package service

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (s *Service) View(identifier int) (*view.View, error) {
	return s.source.View(identifier)
}
