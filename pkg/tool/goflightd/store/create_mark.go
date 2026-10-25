package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"

func (s *Store) CreateMark(v *mark.Mark) error {
	v.Time = v.Time.UTC()

	return s.database.Create(v).Error
}
