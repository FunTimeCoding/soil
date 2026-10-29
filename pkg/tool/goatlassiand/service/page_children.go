package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (s *Service) PageChildren(identifier string) ([]*page.Page, error) {
	return guarded(
		s.confluenceGuard,
		func() ([]*page.Page, error) {
			return s.confluence.ChildPagesByIdentifier(identifier)
		},
	)
}
