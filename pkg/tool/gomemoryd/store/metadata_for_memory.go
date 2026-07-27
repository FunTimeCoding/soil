package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) metadataForMemory(identifier int64) map[string]string {
	rows, e := s.database.Query(
		`SELECT key, value FROM memory_metadata WHERE memory_identifier = ?`,
		identifier,
	)

	if e != nil {
		return nil
	}

	defer errors.LogClose(rows)
	var result map[string]string

	for rows.Next() {
		var key, value string

		if e := rows.Scan(&key, &value); e == nil {
			if result == nil {
				result = map[string]string{}
			}

			result[key] = value
		}
	}

	return result
}
