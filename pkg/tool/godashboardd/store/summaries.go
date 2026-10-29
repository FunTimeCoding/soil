package store

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/model/click"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/types/label_group"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/types/summary"
)

func (s *Store) Summaries() ([]summary.Summary, error) {
	var groups []label_group.Group

	if e := s.mapper.
		Model(click.New()).
		Select("label, count(*) as count, max(id) as newest").
		Group(constant.LabelColumn).
		Order("count DESC").
		Scan(&groups).Error; e != nil {
		return nil, e
	}
	var result []summary.Summary

	for _, g := range groups {
		c := click.New()

		if e := s.mapper.First(c, g.Newest).Error; e != nil {
			return nil, e
		}

		result = append(
			result,
			summary.Summary{Label: g.Label, Count: g.Count, Last: c.CreatedAt},
		)
	}

	return result, nil
}
