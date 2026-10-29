package service

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/plate"
)

func (s *Service) Plate() (*plate.Plate, error) {
	closed, e := s.closedStatus()

	if e != nil {
		return nil, e
	}

	result := plate.New()
	var f error

	if result.Issues, f = s.searchClosed(constant.PlateQuery, closed); f != nil {
		return nil, f
	}

	if result.WatchedIssues, f = s.searchClosed(
		constant.WatchedIssuesQuery,
		closed,
	); f != nil {
		return nil, f
	}

	if result.Newest, f = s.newest(closed); f != nil {
		return nil, f
	}

	if result.Favorites, f = guarded[[]*page.Page](
		s.confluenceGuard,
		s.confluence.Favorites,
	); f != nil {
		return nil, f
	}

	if result.Watched, f = guarded[[]*page.Page](
		s.confluenceGuard,
		s.confluence.Watched,
	); f != nil {
		return nil, f
	}

	return result, nil
}
