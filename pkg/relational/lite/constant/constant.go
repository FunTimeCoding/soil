package constant

const (
	PathEnvironment = "LITE_PATH"
	PathUsage       = "SQLite database path"

	Memory = ":memory:"

	DriverName = "sqlite"

	// Applied per pooled connection - an Exec would only reach one
	FileParameters   = "?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)"
	MemoryParameters = "?_pragma=foreign_keys(1)"
)
