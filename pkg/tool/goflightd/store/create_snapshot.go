package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"

func (s *Store) CreateSnapshot(v snapshot.Snapshot) error {
	return s.database.Create(&v).Error
}
