package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (s *Service) Page(
	identifier string,
	draft bool,
) (*page.Page, error) {
	return guarded(
		s.confluenceGuard,
		func() (*page.Page, error) {
			if draft {
				return s.confluence.DraftPage(identifier)
			}

			return s.confluence.Page(identifier)
		},
	)
}
