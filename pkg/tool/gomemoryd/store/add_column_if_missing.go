package store

import (
	"database/sql"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func addColumnIfMissing(
	d *sql.DB,
	table string,
	column string,
	definition string,
) {
	rows, e := d.Query(fmt.Sprintf("PRAGMA table_info(%s)", table))

	if e != nil {
		return
	}

	defer errors.PanicClose(rows)

	for rows.Next() {
		var cid int
		var name, columnType string
		var notNull int
		var defaultValue *string
		var pk int

		if rows.Scan(
			&cid,
			&name,
			&columnType,
			&notNull,
			&defaultValue,
			&pk,
		) != nil {
			return
		}

		if name == column {
			return
		}
	}

	_, e = d.Exec(
		fmt.Sprintf(
			"ALTER TABLE %s ADD COLUMN %s %s",
			table,
			column,
			definition,
		),
	)
	errors.PanicOnError(e)
}
