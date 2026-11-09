package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue_count"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"maps"
	"slices"
)

func (s *Store) QueuedCounts(
	sessions []session.Session,
) (map[string]int64, error) {
	result := map[string]int64{}

	if len(sessions) == 0 {
		return result, nil
	}

	identifiers := make([]string, 0, len(sessions))
	holders := map[string]string{}

	for _, v := range sessions {
		identifiers = append(identifiers, v.Identifier)

		if v.Callsign != nil && *v.Callsign != "" {
			holders[*v.Callsign] = v.Identifier
		}
	}

	var rows []queue_count.Count
	e := s.database.Model(queue.Stub()).Select(
		"session_identifier, callsign, COUNT(*) AS waiting",
	).Where(
		"consumed_at IS NULL",
	).Where(
		`session_identifier IN ?
		OR ((session_identifier IS NULL OR session_identifier = '')
		AND callsign IN ?)`,
		identifiers,
		slices.Sorted(maps.Keys(holders)),
	).Group(
		"session_identifier, callsign",
	).Scan(
		&rows,
	).Error

	if e != nil {
		return nil, e
	}

	for _, r := range rows {
		owner := r.SessionIdentifier

		if owner == "" {
			owner = holders[r.Callsign]
		}

		if owner != "" {
			result[owner] += r.Waiting
		}
	}

	return result, nil
}
