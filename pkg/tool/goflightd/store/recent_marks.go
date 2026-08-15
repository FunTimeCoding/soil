package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"

func (s *Store) RecentMarks(limit int) ([]mark.Mark, error) {
	var result []mark.Mark
	e := s.database.
		Order("time DESC").
		Limit(limit).
		Find(&result).Error

	return result, e
}
