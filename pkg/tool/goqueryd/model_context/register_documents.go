package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerDocuments() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Push,
			mcp.WithDescription(
				"Push a document into a collection. Chunks, embeds, and indexes in one call. Creates the collection if it doesn't exist.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
			mcp.WithString(
				constant.Path,
				mcp.Required(),
				mcp.Description("Document path within the collection"),
			),
			mcp.WithString(
				constant.Body,
				mcp.Required(),
				mcp.Description("Document content"),
			),
			mcp.WithString(
				constant.SourceType,
				mcp.Description(
					"Source type tag (convenience alias for metadata.source_type)",
				),
			),
			mcp.WithObject(
				constant.Metadata,
				mcp.Description("Key-value metadata pairs for this document"),
			),
		),
		s.push,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Delete,
			mcp.WithDescription(
				"Delete a document from a collection. Removes the document, its embeddings, and orphaned content.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
			mcp.WithString(
				constant.Path,
				mcp.Required(),
				mcp.Description("Document path within the collection"),
			),
		),
		s.delete,
	)
}
