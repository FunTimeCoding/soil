package search_index

import (
	"database/sql"
	"errors"
)

func rollback(t *sql.Tx) {
	e := t.Rollback()

	if e != nil && !errors.Is(e, sql.ErrTxDone) {
		panic(e)
	}
}
