package search_index

import (
	"database/sql"
	"errors"
	library "github.com/funtimecoding/soil/pkg/errors"
)

func (x *Index) position(session string) (int64, string) {
	var consumed int64
	var turn string
	e := x.database.QueryRow(
		"SELECT consumed, turn FROM harbor_file WHERE session = ?",
		session,
	).Scan(&consumed, &turn)

	if errors.Is(e, sql.ErrNoRows) {
		return 0, ""
	}

	library.PanicOnError(e)

	return consumed, turn
}
