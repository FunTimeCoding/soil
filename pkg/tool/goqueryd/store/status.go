package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/status"
)

func (s *Store) Status() (*status.Status, error) {
	r := status.Stub()
	rows, e := s.database.Query(
		`
		SELECT c.name, c.path, c.pattern, c.type,
			COUNT(d.identifier) AS document_count
		FROM collection c
		LEFT JOIN document d ON d.collection = c.name AND d.active = 1
		GROUP BY c.name
	`,
	)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)

	for rows.Next() {
		var c status.CollectionStatus

		if f := rows.Scan(
			&c.Name,
			&c.Path,
			&c.Pattern,
			&c.Type,
			&c.DocumentCount,
		); f != nil {
			return nil, f
		}

		r.Collections = append(r.Collections, c)
		r.TotalDocuments += c.DocumentCount
	}

	if e := rows.Err(); e != nil {
		return nil, e
	}

	row := s.database.QueryRow(
		"SELECT COUNT(*) FROM embedding",
	)

	if e := row.Scan(&r.TotalEmbeddings); e != nil {
		return nil, e
	}

	row = s.database.QueryRow(
		`
		SELECT COUNT(DISTINCT d.hash) FROM document d
		LEFT JOIN embedding e ON e.hash = d.hash
		WHERE d.active = 1 AND e.hash IS NULL
	`,
	)

	if e := row.Scan(&r.PendingEmbeddings); e != nil {
		return nil, e
	}

	return r, nil
}
