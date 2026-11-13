package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreatePage(
	space string,
	parent string,
	title string,
	body string,
	draft bool,
) (*page.Page, error) {
	return peer.Guarded(
		context.Background(),
		s.confluenceGuard,
		func(context.Context) (*page.Page, error) {
			if draft {
				return s.confluence.CreateDraftPage(space, parent, title, body)
			}

			return s.confluence.CreatePage(space, parent, title, body)
		},
	)
}
