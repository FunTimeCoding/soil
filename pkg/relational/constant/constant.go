package constant

const (
	PostgresMessage    = "database: postgres"
	MaxOpenConnections = 10
	MaxIdleConnections = 5

	LitePathEnvironment = "LITE_PATH"
	LitePathUsage       = "SQLite database path"

	LiteMemory = ":memory:"

	LiteMessage = "database: sqlite"

	LiteDriverName  = "sqlite"
	LiteDialectName = "sqlite"

	// Applied per pooled connection - an Exec would only reach one
	LiteFileParameters   = "?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)"
	LiteMemoryParameters = "?_pragma=foreign_keys(1)"

	PostgresAdministratorLocatorEnvironment = "POSTGRES_ADMINISTRATOR_LOCATOR"

	PostgresLocatorEnvironment = "POSTGRES_LOCATOR"
	PostgresLocatorUsage       = "Postgres locator"

	PostgresCommand         = "psql"
	PostgresUserArgument    = "--username"
	PostgresCommandArgument = "--command"
	PostgresFileArgument    = "--file"
	PostgresEchoAllFlag     = "--echo-all"

	PostgresDumpCommand = "pg_dump"

	PostgresDialectName = "postgres"
)
