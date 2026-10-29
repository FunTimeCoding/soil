package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func (s *Store) Authority(name string) (*record.Record, error) {
	var result []record.Record
	e := s.database.
		Where("name = ? AND revoked_at IS NULL", name).
		Limit(1).
		Find(&result).Error

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, not_found.New(constant.AuthorityParameter, name)
	}

	return &result[0], nil
}
