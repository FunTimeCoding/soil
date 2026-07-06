package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func (s *Store) List(f *Filter) ([]entry.Entry, error) {
	q := s.database.Model(entry.New())

	if f != nil {
		if f.System != "" {
			q = q.Where(map[string]any{"system": f.System})
		}

		if f.Service != "" {
			q = q.Where(map[string]any{"service": f.Service})
		}

		if f.User != "" {
			q = q.Where(map[string]any{"user": f.User})
		}

		if !f.Since.IsZero() {
			q = q.Where("timestamp >= ?", f.Since)
		}

		if !f.Until.IsZero() {
			q = q.Where("timestamp <= ?", f.Until)
		}

		if f.Limit > 0 {
			q = q.Limit(f.Limit)
		}
	}

	var result []entry.Entry

	if e := q.Order("timestamp DESC").Find(&result).Error; e != nil {
		return nil, fmt.Errorf("failed to list entries: %w", e)
	}

	return result, nil
}
