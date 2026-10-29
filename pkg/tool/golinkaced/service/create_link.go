package service

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (s *Service) CreateLink(
	l string,
	title string,
	listIdentifier int,
	tags []string,
) (*link.Link, error) {
	return guarded(
		s.guard,
		func() (*link.Link, error) {
			return s.client.CreateLink(l, title, listIdentifier, tags)
		},
	)
}
