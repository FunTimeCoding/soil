package search_index

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/entry"
)

func insert(
	t *sql.Tx,
	n *entry.Entry,
) {
	r, e := t.Exec(
		"INSERT INTO block (identifier, session, turn, role, kind, at) VALUES (?, ?, ?, ?, ?, ?)",
		n.Identifier,
		n.Session,
		n.Turn,
		n.Role,
		n.Kind,
		n.At,
	)
	errors.PanicOnError(e)
	row, e := r.LastInsertId()
	errors.PanicOnError(e)
	_, e = t.Exec(
		"INSERT INTO block_text (rowid, body) VALUES (?, ?)",
		row,
		n.Body,
	)
	errors.PanicOnError(e)
}
