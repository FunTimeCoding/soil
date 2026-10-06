package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerContexts() {
	s.server.AddTool(
		mcp.NewTool(
			constant.AddContext,
			mcp.WithDescription(
				"Add a description for a path prefix within a collection. Attached to search results as hierarchical context.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
			mcp.WithString(
				constant.PathPrefix,
				mcp.Required(),
				mcp.Description("Path prefix to describe"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Required(),
				mcp.Description("Context description"),
			),
		),
		s.addContext,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveContext,
			mcp.WithDescription(
				"Remove context for a path prefix within a collection.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
			mcp.WithString(
				constant.PathPrefix,
				mcp.Required(),
				mcp.Description("Path prefix to remove context for"),
			),
		),
		s.removeContext,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListContexts,
			mcp.WithDescription(
				"List all path prefix contexts across all collections.",
			),
		),
		s.listContexts,
	)
}
