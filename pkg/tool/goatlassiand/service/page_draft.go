package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (s *Service) PageDraft(identifier string) (*page.Page, error) {
	return guarded(
		s.confluenceGuard,
		func() (*page.Page, error) {
			return s.confluence.DraftOverlay(identifier)
		},
	)
}
