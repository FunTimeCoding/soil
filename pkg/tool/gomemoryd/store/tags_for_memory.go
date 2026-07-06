package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) tagsForMemory(identifier int64) []string {
	rows, e := s.database.Query(
		`SELECT tag FROM memory_tag WHERE memory_identifier = ? ORDER BY tag`,
		identifier,
	)

	if e != nil {
		return nil
	}

	defer errors.LogClose(rows)
	var tags []string

	for rows.Next() {
		var tag string

		if e := rows.Scan(&tag); e == nil {
			tags = append(tags, tag)
		}
	}

	return tags
}
