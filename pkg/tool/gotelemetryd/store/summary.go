package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_row"
	"time"
)

func (s *Store) Summary(o *summary_option.Option) ([]summary_row.Row, error) {
	selectClause := "tool, COUNT(*) as count"
	groupClause := "tool"

	switch o.GroupBy {
	case constant.Surface:
		selectClause = "tool, surface, COUNT(*) as count"
		groupClause = "tool, surface"
	case constant.Kind:
		selectClause = "tool, kind, COUNT(*) as count"
		groupClause = "tool, kind"
	case constant.Outcome:
		selectClause = "tool, outcome, COUNT(*) as count"
		groupClause = "tool, outcome"
	}

	query := s.mapper.Model(usage_event.New()).
		Select(selectClause).
		Group(groupClause).
		Order("count DESC")

	if o.Tool != "" {
		query = query.Where("tool = ?", o.Tool)
	}

	if o.Actor != "" {
		query = query.Where("actor = ?", o.Actor)
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

	var result []summary_row.Row

	return result, query.Find(&result).Error
}
