package service

const listSchemasSQL = `
SELECT schema_name, schema_owner
FROM information_schema.schemata
ORDER BY schema_name`

const listTablesSQL = `
SELECT table_name, table_type
FROM information_schema.tables
WHERE table_schema = '%s'
ORDER BY table_name`

const describeTableSQL = `
SELECT
    column_name,
    data_type,
    is_nullable,
    column_default,
    character_maximum_length
FROM information_schema.columns
WHERE table_schema = '%s' AND table_name = '%s'
ORDER BY ordinal_position`

const listIndexesSQL = `
SELECT indexname, indexdef
FROM pg_indexes
WHERE schemaname = '%s' AND tablename = '%s'
ORDER BY indexname`

const tableSizesSQL = `
SELECT
    relname AS table_name,
    n_live_tup AS row_count,
    pg_total_relation_size(c.oid) AS total_bytes,
    pg_size_pretty(pg_total_relation_size(c.oid)) AS total_size
FROM pg_class c
JOIN pg_namespace n ON n.oid = c.relnamespace
JOIN pg_stat_user_tables s ON s.relid = c.oid
WHERE n.nspname = '%s' AND c.relkind = 'r'
ORDER BY pg_total_relation_size(c.oid) DESC`
