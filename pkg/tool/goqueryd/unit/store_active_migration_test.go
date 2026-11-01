package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func TestMigrationDropsTheActiveColumnAndItsRemovedRows(t *testing.T) {
	database := connection.NewMemory()
	defer errors.PanicClose(database)

	for _, s := range []string{
		`CREATE TABLE content (
			hash       TEXT PRIMARY KEY,
			body       TEXT NOT NULL,
			created_at TEXT NOT NULL
		)`,
		`CREATE TABLE document (
			identifier  INTEGER PRIMARY KEY AUTOINCREMENT,
			collection  TEXT NOT NULL,
			path        TEXT NOT NULL,
			title       TEXT NOT NULL,
			hash        TEXT NOT NULL REFERENCES content(hash),
			created_at  TEXT NOT NULL,
			modified_at TEXT NOT NULL,
			active      INTEGER NOT NULL DEFAULT 1,
			UNIQUE(collection, path)
		)`,
		"CREATE INDEX idx_document_collection ON document(collection, active)",
		"CREATE INDEX idx_document_path ON document(path, active)",
		`CREATE VIRTUAL TABLE document_full_text_search USING fts5(
			filepath, title, body,
			tokenize='porter unicode61'
		)`,
		`CREATE TRIGGER document_after_insert
		AFTER INSERT ON document WHEN new.active = 1
		BEGIN
			INSERT INTO document_full_text_search(rowid, filepath, title, body)
			SELECT new.identifier, new.collection || '/' || new.path, new.title,
				(SELECT body FROM content WHERE hash = new.hash);
		END`,
		`CREATE TRIGGER document_after_update
		AFTER UPDATE ON document
		BEGIN
			DELETE FROM document_full_text_search WHERE rowid = old.identifier;
			INSERT INTO document_full_text_search(rowid, filepath, title, body)
			SELECT new.identifier, new.collection || '/' || new.path, new.title,
				(SELECT body FROM content WHERE hash = new.hash)
			WHERE new.active = 1;
		END`,
		`INSERT INTO content VALUES ('a', '# Alfa', 'now'), ('b', '# Bravo', 'now')`,
		`INSERT INTO document
			(collection, path, title, hash, created_at, modified_at, active)
			VALUES
			('test', 'alfa.md', 'Alfa', 'a', 'now', 'now', 1),
			('test', 'bravo.md', 'Bravo', 'b', 'now', 'now', 0)`,
	} {
		_, e := database.Exec(s)
		errors.PanicOnError(e)
	}

	store.New(database, constant.FixtureEmbedModel)
	var columns int
	assert.FatalOnError(
		t,
		database.QueryRow(
			`SELECT COUNT(*) FROM pragma_table_info('document')
			WHERE name = 'active'`,
		).Scan(&columns),
	)
	assert.Integer(t, 0, columns)
	var documents int
	assert.FatalOnError(
		t,
		database.QueryRow("SELECT COUNT(*) FROM document").Scan(&documents),
	)
	assert.Integer(t, 1, documents)
	var triggers int
	assert.FatalOnError(
		t,
		database.QueryRow(
			`SELECT COUNT(*) FROM sqlite_master
			WHERE type IN ('trigger', 'index') AND tbl_name = 'document'
			AND sql LIKE '%active%'`,
		).Scan(&triggers),
	)
	assert.Integer(t, 0, triggers)
}
