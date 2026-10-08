package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/distance"
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
	var candidates []string
	distances := map[string]int{}

	for rows.Next() {
		var path string

		if f := rows.Scan(&path); f != nil {
			return nil, f
		}

		d := distance.Levenshtein(strings.ToLower(path), lower)

		if d <= 5 {
			candidates = append(candidates, path)
			distances[path] = d
		}
	}

	if e := rows.Err(); e != nil {
		return nil, e
	}

	sort.Slice(
		candidates,
		func(i, j int) bool {
			return distances[candidates[i]] < distances[candidates[j]]
		},
	)

	if len(candidates) > limit {
		candidates = candidates[:limit]
	}

	return candidates, nil
}
