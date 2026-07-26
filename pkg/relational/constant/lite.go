package constant

const (
	LitePathEnvironment = "LITE_PATH"
	LitePathUsage       = "SQLite database path"

	LiteMemory = ":memory:"

	LiteMessage = "database: sqlite"

	LiteDriverName  = "sqlite"
	LiteDialectName = "sqlite"

	// Applied per pooled connection - an Exec would only reach one
	LiteFileParameters   = "?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)"
	LiteMemoryParameters = "?_pragma=foreign_keys(1)"
)
