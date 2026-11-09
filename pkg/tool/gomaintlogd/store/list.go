package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/filter"
)

func (s *Store) List(f *filter.Filter) ([]entry.Entry, error) {
	q := s.database.Model(entry.Stub())

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

		if start, started := f.Window.Start(); started {
			q = q.Where("timestamp >= ?", start)
		}

		if end, ended := f.Window.End(); ended {
			q = q.Where("timestamp <= ?", end)
		}

		if f.Limit > 0 {
			q = q.Limit(f.Limit)
		}
	}

	var result []entry.Entry

	if e := q.Order("timestamp DESC").Find(&result).Error; e != nil {
		return nil, fmt.Errorf("list entries: %w", e)
	}

	return result, nil
}
