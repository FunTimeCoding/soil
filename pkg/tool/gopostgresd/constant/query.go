package constant

const (
	ConnectedSessionsQuery = `
SELECT coalesce(application_name, ''), usename,
       coalesce(host(client_addr), 'local'),
       to_char(backend_start AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"')
FROM pg_stat_activity
WHERE datname = current_database()
  AND pid <> pg_backend_pid()
  AND application_name NOT LIKE $1
ORDER BY backend_start`
	DatabaseExistsQuery     = `SELECT 1 FROM pg_database WHERE datname = $1`
	CreateDatabaseStatement = "CREATE DATABASE %s"
	OwnApplicationPattern   = "%s%%"
)

const (
	ListSchemasQuery = `
SELECT schema_name, schema_owner
FROM information_schema.schemata
ORDER BY schema_name`
	ListTablesQuery = `
SELECT table_name, table_type
FROM information_schema.tables
WHERE table_schema = $1
ORDER BY table_name`
	DescribeTableQuery = `
SELECT
    column_name,
    data_type,
    is_nullable,
    column_default,
    character_maximum_length
FROM information_schema.columns
WHERE table_schema = $1 AND table_name = $2
ORDER BY ordinal_position`
	ListIndexesQuery = `
SELECT indexname, indexdef
FROM pg_indexes
WHERE schemaname = $1 AND tablename = $2
ORDER BY indexname`
	TableSizesQuery = `
SELECT
    relname AS table_name,
    n_live_tup AS row_count,
    pg_total_relation_size(c.oid) AS total_bytes,
    pg_size_pretty(pg_total_relation_size(c.oid)) AS total_size
FROM pg_class c
JOIN pg_namespace n ON n.oid = c.relnamespace
JOIN pg_stat_user_tables s ON s.relid = c.oid
WHERE n.nspname = $1 AND c.relkind = 'r'
ORDER BY pg_total_relation_size(c.oid) DESC`
	ExplainPrefix        = "EXPLAIN"
	ExplainAnalyzePrefix = "EXPLAIN ANALYZE"
)
