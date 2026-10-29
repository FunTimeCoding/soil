package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (s *Service) CreatePage(
	space string,
	parent string,
	title string,
	body string,
	draft bool,
) (*page.Page, error) {
	return guarded(
		s.confluenceGuard,
		func() (*page.Page, error) {
			if draft {
				return s.confluence.CreateDraftPage(space, parent, title, body)
			}

			return s.confluence.CreatePage(space, parent, title, body)
		},
	)
}
