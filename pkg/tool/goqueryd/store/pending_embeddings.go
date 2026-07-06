package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) PendingEmbeddings() []PendingDocument {
	rows, e := s.database.Query(
		`
		SELECT d.hash, c.body, MIN(d.path) AS path
		FROM document d
		JOIN content c ON d.hash = c.hash
		LEFT JOIN embedding e ON e.hash = d.hash AND e.sequence = 0
		WHERE d.active = 1 AND e.hash IS NULL
		GROUP BY d.hash
	`,
	)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	var result []PendingDocument

	for rows.Next() {
		var p PendingDocument
		errors.PanicOnError(rows.Scan(&p.Hash, &p.Body, &p.Path))
		result = append(result, p)
	}

	errors.PanicOnError(rows.Err())

	return result
}
