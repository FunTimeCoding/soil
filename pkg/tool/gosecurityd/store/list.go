package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
)

func (s *Store) List(includeMitigated bool) ([]flag.Flag, error) {
	q := s.database.Model(flag.Stub())

	if !includeMitigated {
		q = q.Where(map[string]any{"mitigated": false})
	}

	var result []flag.Flag

	if e := q.Order("created_at DESC").Find(&result).Error; e != nil {
		return nil, fmt.Errorf("list flags: %w", e)
	}

	return result, nil
}
