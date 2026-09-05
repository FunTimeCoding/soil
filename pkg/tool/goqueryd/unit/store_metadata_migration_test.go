package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func TestMigrationWidensExistingMetadataTable(t *testing.T) {
	database := connection.NewMemory()

	for _, s := range []string{
		`CREATE TABLE document (
			identifier  INTEGER PRIMARY KEY AUTOINCREMENT,
			collection  TEXT NOT NULL,
			path        TEXT NOT NULL,
			title       TEXT NOT NULL,
			hash        TEXT NOT NULL,
			created_at  TEXT NOT NULL,
			modified_at TEXT NOT NULL,
			active      INTEGER NOT NULL DEFAULT 1
		)`,
		`CREATE TABLE document_metadata (
			document_identifier INTEGER NOT NULL,
			key   TEXT NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (document_identifier, key)
		)`,
		`INSERT INTO document
			(identifier, collection, path, title, hash, created_at, modified_at)
			VALUES (1, 'test', 'alfa.md', 'Alfa', 'h', 'now', 'now')`,
		`INSERT INTO document_metadata VALUES (1, 'tag', 'build')`,
		`INSERT INTO document_metadata VALUES (1, 'author', 'one')`,
	} {
		_, e := database.Exec(s)
		errors.PanicOnError(e)
	}
	s := store.New(database)
	defer s.Close()
	var key int
	assert.FatalOnError(
		t,
		database.QueryRow(
			`SELECT pk FROM pragma_table_info('document_metadata')
			WHERE name = 'value'`,
		).Scan(&key),
	)
	assert.True(t, key > 0)
	metadata := s.GetMetadata(1)
	assert.Count(t, 2, metadata)
	assert.Count(t, 1, metadata[constant.FixtureTagKey])
	assert.String(t, "build", metadata[constant.FixtureTagKey][0])
	assert.String(t, "one", metadata[constant.FixtureAuthorKey][0])
}
