package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerIndexing() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Index,
			mcp.WithDescription(
				"Re-index filesystem collections. Scans for new, changed, and removed files.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Description(
					"Restrict to a specific collection (omit for all)",
				),
			),
		),
		s.index,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Embed,
			mcp.WithDescription(
				"Generate vector embeddings for documents that are indexed but not yet embedded.",
			),
		),
		s.embed,
	)
}
