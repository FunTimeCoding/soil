package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Page(
	identifier string,
	draft bool,
) (*page.Page, error) {
	return peer.Guarded(
		context.Background(),
		s.confluenceGuard,
		func(context.Context) (*page.Page, error) {
			if draft {
				return s.confluence.DraftPage(identifier)
			}

			return s.confluence.Page(identifier)
		},
	)
}
