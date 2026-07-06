package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) DeleteCollection(name string) bool {
	var exists int
	errors.PanicOnError(
		s.database.QueryRow(
			"SELECT COUNT(*) FROM collection WHERE name = ?",
			name,
		).Scan(&exists),
	)

	if exists == 0 {
		return false
	}

	rows, e := s.database.Query(
		`SELECT DISTINCT d.hash FROM document d
		WHERE d.collection = ?
		AND NOT EXISTS (
			SELECT 1 FROM document o
			WHERE o.hash = d.hash AND o.collection != ?
		)`,
		name,
		name,
	)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	var orphanedHashes []string

	for rows.Next() {
		var hash string
		errors.PanicOnError(rows.Scan(&hash))
		orphanedHashes = append(orphanedHashes, hash)
	}

	errors.PanicOnError(rows.Err())
	_, e = s.database.Exec(
		"DELETE FROM document WHERE collection = ?",
		name,
	)
	errors.PanicOnError(e)

	for _, hash := range orphanedHashes {
		_, e = s.database.Exec("DELETE FROM embedding WHERE hash = ?", hash)
		errors.PanicOnError(e)
		_, e = s.database.Exec("DELETE FROM content WHERE hash = ?", hash)
		errors.PanicOnError(e)
	}

	_, e = s.database.Exec(
		"DELETE FROM context WHERE collection = ?",
		name,
	)
	errors.PanicOnError(e)
	_, e = s.database.Exec(
		"DELETE FROM source_type_tag WHERE collection = ?",
		name,
	)
	errors.PanicOnError(e)
	_, e = s.database.Exec("DELETE FROM collection WHERE name = ?", name)
	errors.PanicOnError(e)

	return true
}
