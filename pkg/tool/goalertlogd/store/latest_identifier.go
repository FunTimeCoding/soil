package store

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Store) LatestIdentifier() (uint, error) {
	var result uint

	return result, s.database.
		Model(record.Stub()).
		Select("COALESCE(MAX(identifier), 0)").
		Scan(&result).Error
}
