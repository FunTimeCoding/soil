package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	atlassianConstant "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SetPageStatus(
	identifier string,
	status string,
) (*page.Page, error) {
	return peer.Guarded(
		context.Background(),
		s.confluenceGuard,
		func(context.Context) (*page.Page, error) {
			var current *page.Page
			var e error

			if status == atlassianConstant.ConfluenceCurrentStatus {
				current, e = s.confluence.DraftOverlay(identifier)
			} else {
				current, e = s.confluence.Page(identifier)
			}

			if e != nil {
				return nil, e
			}

			version := 1

			if status == atlassianConstant.ConfluenceCurrentStatus {
				published, f := s.confluence.Page(identifier)

				if f == nil &&
					published.Raw.Status == atlassianConstant.ConfluenceCurrentStatus {
					version = published.Raw.Version.Number + 1
				}
			}

			return s.confluence.PutPage(
				identifier,
				current.Name,
				current.Raw.Body.Storage.Value,
				version,
				"",
				status,
			)
		},
	)
}
