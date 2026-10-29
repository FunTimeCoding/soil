package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"

func (s *Service) SearchPages(query string) ([]*search_result.Result, error) {
	return guarded(
		s.confluenceGuard,
		func() ([]*search_result.Result, error) {
			return s.confluence.Search(query)
		},
	)
}
