package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pool_name"
)

func (s *Store) poolNames() ([]string, error) {
	var entries []pool_name.PoolName

	if e := s.database.Order(constant.Identifier).Find(&entries).Error; e != nil {
		return nil, e
	}

	if len(entries) == 0 {
		return defaultNames, nil
	}

	result := make([]string, len(entries))

	for i, e := range entries {
		result[i] = e.Name
	}

	return result, nil
}
