package service

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (s *Service) Views() ([]*view.View, error) {
	return s.source.Views()
}
