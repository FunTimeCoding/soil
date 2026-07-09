package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"strings"
)

func initialize(database *sql.DB) {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS collection (
			name    TEXT PRIMARY KEY,
			path    TEXT NOT NULL DEFAULT '',
			pattern TEXT NOT NULL DEFAULT '**/*.md',
			type    TEXT NOT NULL DEFAULT 'filesystem'
		)`,
		`CREATE TABLE IF NOT EXISTS content (
			hash       TEXT PRIMARY KEY,
			body       TEXT NOT NULL,
			created_at TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS document (
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
		"CREATE INDEX IF NOT EXISTS idx_document_collection ON document(collection, active)",
		"CREATE INDEX IF NOT EXISTS idx_document_hash ON document(hash)",
		"CREATE INDEX IF NOT EXISTS idx_document_path ON document(path, active)",
		`CREATE VIRTUAL TABLE IF NOT EXISTS document_full_text_search USING fts5(
			filepath, title, body,
			tokenize='porter unicode61'
		)`,
		`CREATE TRIGGER IF NOT EXISTS document_after_insert
		AFTER INSERT ON document WHEN new.active = 1
		BEGIN
			INSERT INTO document_full_text_search(rowid, filepath, title, body)
			SELECT new.identifier, new.collection || '/' || new.path, new.title,
				(SELECT body FROM content WHERE hash = new.hash);
		END`,
		`CREATE TRIGGER IF NOT EXISTS document_after_delete
		AFTER DELETE ON document
		BEGIN
			DELETE FROM document_full_text_search WHERE rowid = old.identifier;
		END`,
		`CREATE TRIGGER IF NOT EXISTS document_after_update
		AFTER UPDATE ON document
		BEGIN
			DELETE FROM document_full_text_search WHERE rowid = old.identifier;
			INSERT INTO document_full_text_search(rowid, filepath, title, body)
			SELECT new.identifier, new.collection || '/' || new.path, new.title,
				(SELECT body FROM content WHERE hash = new.hash)
			WHERE new.active = 1;
		END`,
		`CREATE TABLE IF NOT EXISTS document_metadata (
			document_identifier INTEGER NOT NULL
				REFERENCES document(identifier) ON DELETE CASCADE,
			key   TEXT NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (document_identifier, key)
		)`,
		"CREATE INDEX IF NOT EXISTS idx_document_metadata_key_value ON document_metadata(key, value)",
		`CREATE TABLE IF NOT EXISTS context (
			collection  TEXT NOT NULL,
			path_prefix TEXT NOT NULL,
			description TEXT NOT NULL,
			PRIMARY KEY (collection, path_prefix)
		)`,
		`CREATE TABLE IF NOT EXISTS embedding (
			hash        TEXT NOT NULL,
			sequence    INTEGER NOT NULL DEFAULT 0,
			position    INTEGER NOT NULL DEFAULT 0,
			vector      BLOB NOT NULL,
			model       TEXT NOT NULL,
			embedded_at TEXT NOT NULL,
			PRIMARY KEY (hash, sequence)
		)`,
		`CREATE TABLE IF NOT EXISTS source_type_tag (
			collection  TEXT NOT NULL DEFAULT '',
			path_prefix TEXT NOT NULL,
			source_type TEXT NOT NULL,
			PRIMARY KEY (collection, path_prefix)
		)`,
	}

	for _, s := range statements {
		_, e := database.Exec(s)
		errors.PanicOnError(e)
	}

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
