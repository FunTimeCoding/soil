package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"sort"
	"strings"
)

func (s *Store) FindSimilarFiles(
	query string,
	limit int,
) ([]string, error) {
	rows, e := s.database.Query(
		"SELECT collection || '/' || path FROM document WHERE active = 1",
	)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	lower := strings.ToLower(query)
	type scored struct {
		path     string
		distance int
	}
	var candidates []scored

	for rows.Next() {
		var path string

		if f := rows.Scan(&path); f != nil {
			return nil, f
		}

		distance := levenshtein(strings.ToLower(path), lower)

		if distance <= 5 {
			candidates = append(
				candidates,
				scored{
					path:     path,
					distance: distance,
				},
			)
		}
	}

	if e := rows.Err(); e != nil {
		return nil, e
	}

	sort.Slice(
		candidates,
		func(i, j int) bool {
			return candidates[i].distance < candidates[j].distance
		},
	)

	if len(candidates) > limit {
		candidates = candidates[:limit]
	}

	result := make([]string, len(candidates))

	for i, c := range candidates {
		result[i] = c.path
	}

	return result, nil
}
