package store

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Store) UnresolvedFingerprints() ([]string, error) {
	var result []string

	return result, s.database.
		Model(record.Stub()).
		Where("ended_at IS NULL").
		Order("fingerprint").
		Pluck("fingerprint", &result).Error
}
