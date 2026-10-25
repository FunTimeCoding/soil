package constant

const (
	LiteTableExistsQuery   = "SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name = ?"
	LiteTableQuery         = "SELECT name FROM pragma_table_list WHERE schema = 'main' AND type = 'table' AND name NOT LIKE 'sqlite_%'"
	LiteTextColumnQuery    = "SELECT name FROM pragma_table_info(?) WHERE type = '' OR upper(type) GLOB '*TEXT*' OR upper(type) GLOB '*CHAR*' OR upper(type) GLOB '*CLOB*' OR upper(type) GLOB '*DATE*' OR upper(type) GLOB '*TIME*'"
	LiteUniversalQuery     = `UPDATE "%[1]s" SET "%[2]s" = strftime('%%Y-%%m-%%d %%H:%%M:%%S', "%[2]s") || substr("%[2]s", 20, length("%[2]s") - 25) || '+00:00' WHERE typeof("%[2]s") = 'text' AND length("%[2]s") BETWEEN 25 AND 35 AND (length("%[2]s") = 25 OR substr("%[2]s", 20, 1) = '.') AND "%[2]s" GLOB '[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]*[+-][0-9][0-9]:[0-9][0-9]' AND substr("%[2]s", -6) <> '+00:00' AND strftime('%%s', "%[2]s") IS NOT NULL`
	LiteUniversalMarkQuery = `CREATE TABLE "%s" (normalized_at TEXT NOT NULL)`
)
