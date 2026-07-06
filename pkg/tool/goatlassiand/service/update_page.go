package service

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (s *Service) UpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) (*page.Page, error) {
	current, e := s.confluence.Page(identifier)

	if e != nil {
		return nil, e
	}

	return s.confluence.PutPage(
		identifier,
		title,
		page.ToStorage(markdown),
		current.Raw.Version.Number+1,
		message,
		constant.CurrentStatus,
	)
}
