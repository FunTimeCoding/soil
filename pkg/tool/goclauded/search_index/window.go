package search_index

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/block"
)

func (x *Index) Window(
	session string,
	around string,
	count int,
) ([]*block.Block, error) {
	rows, e := x.database.Query(
		`WITH ordered AS (
			SELECT rowid, identifier, role, kind, at,
				row_number() OVER (ORDER BY rowid) AS position
			FROM block WHERE session = ?
		),
		anchor AS (
			SELECT min(position) AS position FROM ordered WHERE identifier = ?
		)
		SELECT o.identifier, o.role, o.kind, o.at, t.body
		FROM ordered o
		JOIN anchor a
		JOIN block_text t ON t.rowid = o.rowid
		WHERE o.position BETWEEN a.position - ? AND a.position + ?
		ORDER BY o.position`,
		session,
		around,
		count,
		count,
	)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var result []*block.Block

	for rows.Next() {
		var identifier, role, kind, at, body string

		if f := rows.Scan(&identifier, &role, &kind, &at, &body); f != nil {
			return nil, f
		}

		result = append(result, block.New(identifier, role, kind, at, body))
	}

	return result, rows.Err()
}
