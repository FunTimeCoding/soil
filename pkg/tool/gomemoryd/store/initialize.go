package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
)

func initialize(database *sql.DB) {
	statements := []string{
		"PRAGMA journal_mode = WAL",
		"PRAGMA foreign_keys = ON",
		`CREATE TABLE IF NOT EXISTS memory (
			identifier  INTEGER PRIMARY KEY AUTOINCREMENT,
			content     TEXT NOT NULL,
			description TEXT NOT NULL,
			type        TEXT NOT NULL,
			created_at  TEXT NOT NULL,
			updated_at  TEXT NOT NULL,
			embedding   BLOB,
			is_active   INTEGER NOT NULL DEFAULT 1,
			parent_identifier INTEGER REFERENCES memory(identifier)
		)`,
		"CREATE INDEX IF NOT EXISTS idx_memory_type ON memory(type, is_active)",
		"CREATE INDEX IF NOT EXISTS idx_memory_active ON memory(is_active)",
		`CREATE TABLE IF NOT EXISTS memory_version (
			identifier  INTEGER PRIMARY KEY AUTOINCREMENT,
			memory_identifier   INTEGER NOT NULL REFERENCES memory(identifier),
			content     TEXT NOT NULL,
			description TEXT NOT NULL,
			changed_at  TEXT NOT NULL,
			change_type TEXT NOT NULL
		)`,
		"CREATE INDEX IF NOT EXISTS idx_memory_version_memory ON memory_version(memory_identifier)",
		`CREATE TABLE IF NOT EXISTS memory_tag (
			memory_identifier INTEGER NOT NULL REFERENCES memory(identifier),
			tag       TEXT NOT NULL,
			PRIMARY KEY (memory_identifier, tag)
		)`,
		"CREATE INDEX IF NOT EXISTS idx_memory_tag_tag ON memory_tag(tag)",
		`CREATE TABLE IF NOT EXISTS memory_relation (
			source_identifier  INTEGER NOT NULL REFERENCES memory(identifier),
			target_identifier  INTEGER NOT NULL REFERENCES memory(identifier),
			created_at TEXT NOT NULL,
			PRIMARY KEY (source_identifier, target_identifier)
		)`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS memory_full_text_search USING fts5(
			content, description,
			tokenize='porter unicode61'
		)`,
		`CREATE TRIGGER IF NOT EXISTS memory_full_text_search_insert
		AFTER INSERT ON memory WHEN new.is_active = 1
		BEGIN
			INSERT INTO memory_full_text_search(rowid, content, description)
			VALUES (new.identifier, new.content, new.description);
		END`,
		`CREATE TRIGGER IF NOT EXISTS memory_full_text_search_delete
		AFTER DELETE ON memory
		BEGIN
			DELETE FROM memory_full_text_search WHERE rowid = old.identifier;
		END`,
		`CREATE TRIGGER IF NOT EXISTS memory_full_text_search_update
		AFTER UPDATE ON memory
		BEGIN
			DELETE FROM memory_full_text_search WHERE rowid = old.identifier;
			INSERT INTO memory_full_text_search(rowid, content, description)
			SELECT new.identifier, new.content, new.description
			WHERE new.is_active = 1;
		END`,
		`CREATE TABLE IF NOT EXISTS impression (
			identifier  INTEGER PRIMARY KEY AUTOINCREMENT,
			content    TEXT NOT NULL,
			source     TEXT NOT NULL DEFAULT '',
			created_at TEXT NOT NULL
		)`,
	}

	for _, s := range statements {
		_, e := database.Exec(s)
		errors.PanicOnError(e)
	}

	migrate(database)
}
