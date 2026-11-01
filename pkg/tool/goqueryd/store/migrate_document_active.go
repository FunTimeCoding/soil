package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
)

func migrateDocumentActive(database *sql.DB) {
	var name string

	if database.QueryRow(
		`SELECT name FROM pragma_table_info('document') WHERE name = 'active'`,
	).Scan(&name) != nil {
		return
	}

	for _, s := range []string{
		"DELETE FROM document WHERE active = 0",
		"DROP TRIGGER IF EXISTS document_after_insert",
		"DROP TRIGGER IF EXISTS document_after_update",
		"DROP INDEX IF EXISTS idx_document_collection",
		"DROP INDEX IF EXISTS idx_document_path",
		"ALTER TABLE document DROP COLUMN active",
	} {
		_, e := database.Exec(s)
		errors.PanicOnError(e)
	}
}
