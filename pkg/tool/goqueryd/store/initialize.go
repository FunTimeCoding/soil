package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"slices"
	"strings"
)

func initialize(database *sql.DB) {
	migrateDocumentActive(database)
	statements := slices.Concat(documentStatements(), embeddingStatements())

	for _, s := range statements {
		_, e := database.Exec(s)
		errors.PanicOnError(e)
	}

	migrateMetadataValue(database)
	migrations := []string{
		"ALTER TABLE collection ADD COLUMN type TEXT NOT NULL DEFAULT 'filesystem'",
	}

	for _, m := range migrations {
		_, e := database.Exec(m)

		if e != nil && !strings.Contains(e.Error(), "duplicate column") {
			errors.PanicOnError(e)
		}
	}
}
