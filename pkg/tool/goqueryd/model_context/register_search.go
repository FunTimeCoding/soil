package model_context

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSearch() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Search,
			mcp.WithDescription(
				"Search indexed documents using hybrid BM25 keyword + vector similarity + cross-encoder reranking. Returns ranked results with snippets.",
			),
			mcp.WithString(
				generative.ParameterQuery,
				mcp.Required(),
				mcp.Description("Search query"),
			),
			mcp.WithNumber(
				generative.ParameterLimit,
				mcp.Description("Maximum number of results (default 10)"),
			),
			mcp.WithString(
				constant.Collection,
				mcp.Description("Restrict search to a specific collection"),
			),
			mcp.WithString(
				constant.Mode,
				mcp.Description("Search mode: hybrid (default) or keyword"),
			),
			mcp.WithBoolean(
				constant.Full,
				mcp.Description(
					"Include full document body in results (default false)",
				),
			),
			mcp.WithString(
				constant.SourceType,
				mcp.Description("Filter results by source type"),
			),
			mcp.WithObject(
				constant.Metadata,
				mcp.Description(
					"Filter results by metadata key-value pairs (exact match)",
				),
			),
		),
		s.search,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Status,
			mcp.WithDescription(
				"Show index status: total documents, embeddings, pending embeddings, and per-collection statistics.",
			),
		),
		s.status,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Get,
			mcp.WithDescription(
				"Get a single document by path. Accepts virtual paths (qmd://collection/path) or relative paths (collection/path). Returns full document content with title, context, and metadata.",
			),
			mcp.WithString(
				constant.Path,
				mcp.Required(),
				mcp.Description(
					"Document path (qmd://collection/path or collection/path)",
				),
			),
		),
		s.get,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.List,
			mcp.WithDescription(
				"List documents in a collection, ordered by recency. Optional metadata filter, pagination, and full body. Response includes facets showing metadata key distribution.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
			mcp.WithString(
				constant.SourceType,
				mcp.Description("Filter by source type"),
			),
			mcp.WithObject(
				constant.Metadata,
				mcp.Description(
					"Filter results by metadata key-value pairs (exact match)",
				),
			),
			mcp.WithNumber(
				generative.ParameterLimit,
				mcp.Description(
					"Maximum number of results (default 10, 0 for all)",
				),
			),
			mcp.WithNumber(
				generative.ParameterOffset,
				mcp.Description("Number of results to skip (default 0)"),
			),
			mcp.WithBoolean(
				constant.Full,
				mcp.Description("Include full document body (default false)"),
			),
		),
		s.list,
	)
}
