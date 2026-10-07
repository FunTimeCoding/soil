package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"time"
)

func (s *Store) Recent(o *query_option.Option) ([]usage_event.Event, error) {
	query := s.mapper.Order("created_at DESC")

	if o.Tool != "" {
		query = query.Where("tool = ?", o.Tool)
	}

	if o.Surface != "" {
		query = query.Where("surface = ?", o.Surface)
	}

	if o.Actor != "" {
		query = query.Where("actor = ?", o.Actor)
	}

	if o.Kind != "" {
		query = query.Where("kind = ?", o.Kind)
	}

	if o.Since != "" {
		t, e := time.Parse(time.RFC3339, o.Since)

		if e == nil {
			query = query.Where("created_at >= ?", t)
		}
	}

	if o.Until != "" {
		t, e := time.Parse(time.RFC3339, o.Until)

		if e == nil {
			query = query.Where("created_at <= ?", t)
		}
	}

	if o.Offset > 0 {
		query = query.Offset(o.Offset)
	}

	if o.Limit > 0 {
		query = query.Limit(o.Limit)
	}

	var result []usage_event.Event

	return result, query.Find(&result).Error
}
