package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"

func (s *Store) SnapshotCount() (int64, error) {
	var result int64
	e := s.database.Model(snapshot.Stub()).Count(&result).Error

	return result, e
}
