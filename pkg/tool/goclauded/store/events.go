package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
)

func (s *Store) Events(q *event_query.Query) ([]event.Event, error) {
	g := s.database.Model(event.Stub())

	if !q.Since().IsZero() {
		g = g.Where("created_at >= ?", q.Since())
	}

	if !q.Before().IsZero() {
		g = g.Where("created_at <= ?", q.Before())
	}

	if len(q.Kinds()) == 1 {
		g = g.Where("kind = ?", q.Kinds()[0])
	} else if len(q.Kinds()) > 1 {
		g = g.Where("kind IN ?", q.Kinds())
	}

	var result []event.Event

	if e := g.Order(
		"created_at DESC",
	).Limit(q.Limit()).Offset(q.Offset()).Find(&result).Error; e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return result, nil
	}

	var identifiers []uint

	for _, v := range result {
		identifiers = append(identifiers, v.Identifier)
	}

	metadata := s.EventMetadataByEvents(identifiers)

	for i := range result {
		result[i].Metadata = metadata[result[i].Identifier]
	}

	return result, nil
}
