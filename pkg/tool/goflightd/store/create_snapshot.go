package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"

func (s *Store) CreateSnapshot(v snapshot.Snapshot) error {
	v.Time = v.Time.UTC()

	return s.database.Create(&v).Error
}
