package model_context

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerCollections() {
	s.server.AddTool(
		mcp.NewTool(
			constant.AddCollection,
			mcp.WithDescription(
				"Register a filesystem collection for indexing.",
			),
			mcp.WithString(
				generative.ParameterName,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
			mcp.WithString(
				constant.Path,
				mcp.Required(),
				mcp.Description("Filesystem path to index"),
			),
			mcp.WithString(
				constant.Pattern,
				mcp.Description("Glob pattern for files (default **/*.md)"),
			),
		),
		s.addCollection,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteCollection,
			mcp.WithDescription(
				"Delete a collection and all its documents, embeddings, contexts, and source type tags.",
			),
			mcp.WithString(
				generative.ParameterName,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
		),
		s.deleteCollection,
	)
}
