package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/fable_snapshot"

func (s *Store) LatestFableSnapshot() (*fable_snapshot.Snapshot, error) {
	var result fable_snapshot.Snapshot
	r := s.database.Order("created_at DESC").Limit(1).Find(&result)

	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, nil
	}

	return &result, nil
}
