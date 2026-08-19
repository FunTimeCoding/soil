package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
)

func migrateMetadataValue(database *sql.DB) {
	var key int

	if database.QueryRow(
		`SELECT pk FROM pragma_table_info('document_metadata')
		WHERE name = 'value'`,
	).Scan(&key) != nil {
		return
	}

	if key > 0 {
		return
	}

	for _, s := range []string{
		`CREATE TABLE document_metadata_new (
			document_identifier INTEGER NOT NULL
				REFERENCES document(identifier) ON DELETE CASCADE,
			key   TEXT NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (document_identifier, key, value)
		)`,
		`INSERT INTO document_metadata_new (document_identifier, key, value)
			SELECT document_identifier, key, value FROM document_metadata`,
		"DROP TABLE document_metadata",
		"ALTER TABLE document_metadata_new RENAME TO document_metadata",
		`CREATE INDEX IF NOT EXISTS idx_document_metadata_key_value
			ON document_metadata(key, value)`,
	} {
		_, e := database.Exec(s)
		errors.PanicOnError(e)
	}
}
