package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Store) allEmbeddings(collection string) []embeddingCandidate {
	sql := `
		SELECT
			d.collection || '/' || d.path AS filepath,
			d.collection,
			d.path,
			d.title,
			d.hash,
			content.body,
			e.position,
			e.vector
		FROM embedding e
		JOIN document d ON d.hash = e.hash AND d.active = 1
		JOIN content ON content.hash = d.hash`
	var arguments []any

	if collection != "" {
		sql = join.Empty(sql, " WHERE d.collection = ?")
		arguments = append(arguments, collection)
	}

	rows, e := s.database.Query(sql, arguments...)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	var result []embeddingCandidate

	for rows.Next() {
		var c embeddingCandidate
		var blob []byte
		errors.PanicOnError(
			rows.Scan(
				&c.filePath,
				&c.collection,
				&c.path,
				&c.title,
				&c.hash,
				&c.body,
				&c.position,
				&blob,
			),
		)
		c.vector = bytesToFloat32(blob)
		result = append(result, c)
	}

	errors.PanicOnError(rows.Err())

	return result
}
