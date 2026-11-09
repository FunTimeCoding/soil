package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.instances.Register(s.server)
	s.server.AddTool(
		mcp.NewTool(
			constant.Query,
			mcp.WithDescription("Execute SQL against the active instance"),
			mcp.WithString(
				"sql",
				mcp.Required(),
				mcp.Description("SQL query to execute"),
			),
		),
		mcp.NewTypedToolHandler(s.query),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Explain,
			mcp.WithDescription("Show the execution plan for a SQL query"),
			mcp.WithString(
				"sql",
				mcp.Required(),
				mcp.Description("SQL query to explain"),
			),
			mcp.WithBoolean(
				"analyze",
				mcp.Description(
					"Run EXPLAIN ANALYZE for actual execution stats",
				),
			),
		),
		mcp.NewTypedToolHandler(s.explain),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSchemas,
			mcp.WithDescription("List schemas in the active database"),
		),
		mcp.NewTypedToolHandler(s.listSchemas),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTables,
			mcp.WithDescription("List tables in a schema"),
			mcp.WithString(
				"schema",
				mcp.Description("Schema name (default: public)"),
			),
		),
		mcp.NewTypedToolHandler(s.listTables),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DescribeTable,
			mcp.WithDescription(
				"Show columns, types, and constraints for a table",
			),
			mcp.WithString(
				"table",
				mcp.Required(),
				mcp.Description("Table name"),
			),
			mcp.WithString(
				"schema",
				mcp.Description("Schema name (default: public)"),
			),
		),
		mcp.NewTypedToolHandler(s.describeTable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListIndexes,
			mcp.WithDescription("List indexes on a table"),
			mcp.WithString(
				"table",
				mcp.Required(),
				mcp.Description("Table name"),
			),
			mcp.WithString(
				"schema",
				mcp.Description("Schema name (default: public)"),
			),
		),
		mcp.NewTypedToolHandler(s.listIndexes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.TableSizes,
			mcp.WithDescription("Show row counts and disk usage for tables"),
			mcp.WithString(
				"schema",
				mcp.Description("Schema name (default: public)"),
			),
		),
		mcp.NewTypedToolHandler(s.tableSizes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateBackup,
			mcp.WithDescription(
				"Dump the active instance's database to a new backup with pg_dump; a large database takes minutes",
			),
		),
		mcp.NewTypedToolHandler(s.createBackup),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListBackups,
			mcp.WithDescription(
				"List the active instance's backups, newest first",
			),
		),
		mcp.NewTypedToolHandler(s.listBackups),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RestoreBackup,
			mcp.WithDescription(
				"Restore a backup of the active instance into a new database, or in place when the target is the instance's own database and no other session is connected",
			),
			mcp.WithString(
				constant.BackupParameter,
				mcp.Required(),
				mcp.Description("Backup name, as list_backups shows it"),
			),
			mcp.WithString(
				"target",
				mcp.Description(
					"Database to restore into (default: a new <database>_restore_<backup>)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.restoreBackup),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteBackup,
			mcp.WithDescription("Delete a backup of the active instance"),
			mcp.WithString(
				constant.BackupParameter,
				mcp.Required(),
				mcp.Description("Backup name, as list_backups shows it"),
			),
		),
		mcp.NewTypedToolHandler(s.deleteBackup),
	)
}
