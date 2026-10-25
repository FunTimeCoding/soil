package zone

import (
	"database/sql"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
)

func Normalize(d *sql.DB) {
	var count int
	errors.PanicOnError(
		d.QueryRow(
			constant.LiteTableExistsQuery,
			constant.LiteUniversalTable,
		).Scan(&count),
	)

	if count > 0 {
		return
	}

	t, e := d.Begin()
	errors.PanicOnError(e)
	defer rollback(t)

	for _, table := range names(t, constant.LiteTableQuery) {
		for _, column := range names(t, constant.LiteTextColumnQuery, table) {
			_, f := t.Exec(
				fmt.Sprintf(constant.LiteUniversalQuery, table, column),
			)
			errors.PanicOnError(f)
		}
	}

	_, e = t.Exec(
		fmt.Sprintf(
			constant.LiteUniversalMarkQuery,
			constant.LiteUniversalTable,
		),
	)
	errors.PanicOnError(e)
	errors.PanicOnError(t.Commit())
}
