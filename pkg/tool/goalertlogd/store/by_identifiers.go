package store

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Store) ByIdentifiers(identifiers []uint) ([]record.Record, error) {
	var result []record.Record

	if len(identifiers) == 0 {
		return result, nil
	}

	return result, s.database.
		Where("identifier IN ?", identifiers).
		Order("identifier").
		Find(&result).Error
}
