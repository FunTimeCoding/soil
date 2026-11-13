package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	atlassianConstant "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) EditPage(
	identifier string,
	oldText string,
	newText string,
	title string,
	message string,
	draft bool,
) (*page.Page, error) {
	current, e := s.Page(identifier, draft)

	if e != nil {
		return nil, e
	}

	markdown := page.ToMarkdown(current.Raw.Body.Storage.Value)
	newMarkdown, f := ReplaceUnique(markdown, oldText, newText)

	if f != nil {
		return nil, f
	}

	if title == "" {
		title = current.Name
	}

	status := atlassianConstant.ConfluenceCurrentStatus

	if draft {
		status = atlassianConstant.ConfluenceDraftStatus
	}

	return peer.Guarded(
		context.Background(),
		s.confluenceGuard,
		func(context.Context) (*page.Page, error) {
			return s.confluence.PutPage(
				identifier,
				title,
				page.ToStorage(newMarkdown),
				current.Raw.Version.Number+1,
				message,
				status,
			)
		},
	)
}
