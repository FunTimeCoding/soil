package store

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/model/click"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/types/summary"
)

func (s *Store) Summaries() ([]summary.Summary, error) {
	var result []summary.Summary

	return result, s.mapper.
		Model(click.New()).
		Select("label, count(*) as count, max(created_at) as last").
		Group(constant.LabelColumn).
		Order("count DESC").
		Scan(&result).Error
}
