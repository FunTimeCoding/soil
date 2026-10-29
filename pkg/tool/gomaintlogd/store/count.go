package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func (s *Store) Count() (int, error) {
	var result int64

	if e := s.database.Model(entry.New()).Count(&result).Error; e != nil {
		return 0, fmt.Errorf("count entries: %w", e)
	}

	return int(result), nil
}
