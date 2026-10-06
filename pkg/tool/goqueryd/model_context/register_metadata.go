package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMetadata() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMetadata,
			mcp.WithDescription(
				"List distinct metadata keys and value distributions for a collection. Shows which fields exist and what values they take - use before filtering to discover available metadata.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
			mcp.WithString(
				constant.Key,
				mcp.Description(
					"Specific key to expand (shows all values regardless of cardinality)",
				),
			),
		),
		s.listMetadata,
	)
}
