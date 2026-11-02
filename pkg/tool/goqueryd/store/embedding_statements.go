package store

func embeddingStatements() []string {
	return []string{
		`CREATE TABLE IF NOT EXISTS document_metadata (
			document_identifier INTEGER NOT NULL
				REFERENCES document(identifier) ON DELETE CASCADE,
			key   TEXT NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (document_identifier, key, value)
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
		`CREATE TABLE IF NOT EXISTS embedding_failure (
			hash      TEXT NOT NULL,
			sequence  INTEGER NOT NULL,
			position  INTEGER NOT NULL,
			reason    TEXT NOT NULL,
			failed_at TEXT NOT NULL,
			PRIMARY KEY (hash, sequence)
		)`,
		`CREATE TABLE IF NOT EXISTS source_type_tag (
			collection  TEXT NOT NULL DEFAULT '',
			path_prefix TEXT NOT NULL,
			source_type TEXT NOT NULL,
			PRIMARY KEY (collection, path_prefix)
		)`,
	}
}
