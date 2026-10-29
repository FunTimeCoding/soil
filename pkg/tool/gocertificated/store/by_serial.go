package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func (s *Store) BySerial(serial string) (*record.Record, error) {
	var result []record.Record
	e := s.database.
		Where("serial = ?", serial).
		Limit(1).
		Find(&result).Error

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, not_found.New(constant.SerialParameter, serial)
	}

	return &result[0], nil
}
