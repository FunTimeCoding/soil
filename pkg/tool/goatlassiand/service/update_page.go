package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) (*page.Page, error) {
	return peer.Guarded(
		context.Background(),
		s.confluenceGuard,
		func(context.Context) (*page.Page, error) {
			return s.confluence.UpdatePage(identifier, title, markdown, message)
		},
	)
}
