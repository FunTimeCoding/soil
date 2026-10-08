package search_index

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/block"
)

func (x *Index) ConversationBlocks(
	session string,
	kinds []string,
) ([]*block.Block, error) {
	kinds = kindsOrDefault(kinds)
	arguments := []any{session}

	for _, kind := range kinds {
		arguments = append(arguments, kind)
	}

	rows, e := x.database.Query(
		fmt.Sprintf(
			`SELECT b.identifier, b.role, b.kind, b.at, t.body
			FROM block b JOIN block_text t ON t.rowid = b.rowid
			WHERE b.session = ? AND b.kind IN (%s)
			ORDER BY b.rowid`,
			placeholders(len(kinds)),
		),
		arguments...,
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
