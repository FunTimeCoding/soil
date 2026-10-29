package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (s *Service) UpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) (*page.Page, error) {
	return guarded(
		s.confluenceGuard,
		func() (*page.Page, error) {
			return s.confluence.UpdatePage(identifier, title, markdown, message)
		},
	)
}
