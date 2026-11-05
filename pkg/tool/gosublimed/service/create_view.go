package service

import "github.com/funtimecoding/soil/pkg/sublime/view"

func (s *Service) CreateView(
	title string,
	content string,
	syntax string,
) (*view.View, error) {
	return guarded(
		s.guard,
		func() (*view.View, error) {
			return s.client.CreateView(title, content, syntax)
		},
	)
}
