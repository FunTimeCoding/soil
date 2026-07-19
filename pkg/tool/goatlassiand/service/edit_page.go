package service

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (s *Service) EditPage(
	identifier string,
	oldText string,
	newText string,
	title string,
	message string,
	draft bool,
) (*page.Page, error) {
	var current *page.Page
	var e error

	if draft {
		current, e = s.confluence.DraftPage(identifier)
	} else {
		current, e = s.confluence.Page(identifier)
	}

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

	status := constant.CurrentStatus

	if draft {
		status = constant.DraftStatus
	}

	return s.confluence.PutPage(
		identifier,
		title,
		page.ToStorage(newMarkdown),
		current.Raw.Version.Number+1,
		message,
		status,
	)
}
