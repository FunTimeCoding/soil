package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/impression_outbox"
)

func (s *Store) PendingImpressions() ([]*impression_outbox.Impression, error) {
	var result []*impression_outbox.Impression

	if e := s.database.Order(constant.Identifier).Find(&result).Error; e != nil {
		return nil, e
	}

	return result, nil
}
