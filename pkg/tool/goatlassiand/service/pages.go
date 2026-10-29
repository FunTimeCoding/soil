package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (s *Service) Pages(
	space string,
	status string,
) ([]*page.Page, error) {
	return guarded(
		s.confluenceGuard,
		func() ([]*page.Page, error) {
			return s.confluence.PagesBySpace(space, status)
		},
	)
}
