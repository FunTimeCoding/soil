package store

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Store) BySerial(serial string) (*record.Record, error) {
	var result []record.Record
	e := s.database.
		Where("serial = ?", serial).
		Limit(1).
		Find(&result).Error

	if e != nil || len(result) == 0 {
		return nil, e
	}

	return &result[0], nil
}
