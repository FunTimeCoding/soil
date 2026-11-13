package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SearchPages(query string) ([]*search_result.Result, error) {
	return peer.Guarded(
		context.Background(),
		s.confluenceGuard,
		func(context.Context) ([]*search_result.Result, error) {
			return s.confluence.Search(query)
		},
	)
}
