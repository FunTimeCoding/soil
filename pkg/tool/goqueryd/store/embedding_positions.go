package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) EmbeddingPositions() (map[string]map[int]int, error) {
	rows, e := s.database.Query(
		"SELECT hash, sequence, position FROM embedding",
	)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	result := map[string]map[int]int{}

	for rows.Next() {
		var hash string
		var sequence, position int

		if f := rows.Scan(&hash, &sequence, &position); f != nil {
			return nil, f
		}

		if result[hash] == nil {
			result[hash] = map[int]int{}
		}

		result[hash][sequence] = position
	}

	return result, rows.Err()
}
