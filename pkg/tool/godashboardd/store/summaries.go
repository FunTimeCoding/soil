package store

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"

func (s *Store) Summaries() ([]Summary, error) {
	var result []Summary

	return result, s.mapper.
		Model(NewClick()).
		Select("label, count(*) as count, max(created_at) as last").
		Group(constant.LabelColumn).
		Order("count DESC").
		Scan(&result).Error
}
