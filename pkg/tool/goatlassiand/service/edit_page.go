package service

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
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

	status := constant.ConfluenceCurrentStatus

	if draft {
		status = constant.ConfluenceDraftStatus
	}

	return guarded(
		s.confluenceGuard,
		func() (*page.Page, error) {
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
