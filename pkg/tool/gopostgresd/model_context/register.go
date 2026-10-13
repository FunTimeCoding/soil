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
}
