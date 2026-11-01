package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/dashboard_view"
	"time"
)

func (s *Service) DashboardView() (*dashboard_view.View, error) {
	result := dashboard_view.New()

	if e := s.countMemories(result); e != nil {
		return nil, e
	}

	relations, e := s.VisibleRelations()

	if e != nil {
		return nil, e
	}

	result.Relations = len(relations)

	for _, v := range relations {
		if v.Type == "" {
			result.Untyped++
		}
	}

	now := time.Now()
	result.Versions, e = s.RecentVersions(
		since(now, constant.DashboardVersionWindow),
		constant.DashboardVersionLimit,
	)

	if e != nil {
		return nil, e
	}

	result.Impressions, e = s.store.RecentImpressions(
		since(now, constant.DashboardImpressionWindow),
	)

	if e != nil {
		return nil, e
	}

	return result, nil
}
