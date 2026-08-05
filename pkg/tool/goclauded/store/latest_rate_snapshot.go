package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/rate_snapshot"

func (s *Store) LatestRateSnapshot() (*rate_snapshot.Snapshot, error) {
	var result rate_snapshot.Snapshot
	r := s.database.Order("created_at DESC").Limit(1).Find(&result)

	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, nil
	}

	return &result, nil
}
