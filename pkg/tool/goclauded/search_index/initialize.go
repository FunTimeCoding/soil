package search_index

import (
	"database/sql"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func initialize(d *sql.DB) {
	var version int
	errors.PanicOnError(d.QueryRow("PRAGMA user_version").Scan(&version))

	if version != constant.SearchIndexVersion {
		for _, table := range []string{"block", "block_text", "harbor_file"} {
			_, e := d.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
			errors.PanicOnError(e)
		}
	}

	for _, statement := range []string{
		`CREATE TABLE IF NOT EXISTS block (
			rowid      INTEGER PRIMARY KEY,
			identifier TEXT NOT NULL,
			session    TEXT NOT NULL,
			turn       TEXT NOT NULL,
			role       TEXT NOT NULL,
			kind       TEXT NOT NULL,
			at         TEXT NOT NULL
		)`,
		"CREATE INDEX IF NOT EXISTS block_session ON block(session)",
		`CREATE VIRTUAL TABLE IF NOT EXISTS block_text USING fts5(
			body,
			tokenize='trigram'
		)`,
		`CREATE TABLE IF NOT EXISTS harbor_file (
			session  TEXT PRIMARY KEY,
			consumed INTEGER NOT NULL,
			turn     TEXT NOT NULL
		)`,
		fmt.Sprintf("PRAGMA user_version = %d", constant.SearchIndexVersion),
	} {
		_, e := d.Exec(statement)
		errors.PanicOnError(e)
	}
}
