package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Search,
			mcp.WithDescription(
				"Search indexed documents using hybrid BM25 keyword + vector similarity + cross-encoder reranking. Returns ranked results with snippets.",
			),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("Search query"),
			),
			mcp.WithNumber(
				parameter.Limit,
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
				mcp.Description("Include full document body in results (default false)"),
			),
			mcp.WithString(
				constant.SourceType,
				mcp.Description("Filter results by source type"),
			),
			mcp.WithObject(
				constant.Metadata,
				mcp.Description("Filter results by metadata key-value pairs (exact match)"),
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
				mcp.Description("Document path (qmd://collection/path or collection/path)"),
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
				mcp.Description("Filter results by metadata key-value pairs (exact match)"),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Description("Maximum number of results (default 10, 0 for all)"),
			),
			mcp.WithNumber(
				parameter.Offset,
				mcp.Description("Number of results to skip (default 0)"),
			),
			mcp.WithBoolean(
				constant.Full,
				mcp.Description("Include full document body (default false)"),
			),
		),
		s.list,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Index,
			mcp.WithDescription(
				"Re-index filesystem collections. Scans for new, changed, and removed files.",
			),
			mcp.WithString(
				constant.Collection,
				mcp.Description("Restrict to a specific collection (omit for all)"),
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
	s.server.AddTool(
		mcp.NewTool(
			constant.AddCollection,
			mcp.WithDescription(
				"Register a filesystem collection for indexing.",
			),
			mcp.WithString(
				parameter.Name,
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
				parameter.Name,
				mcp.Required(),
				mcp.Description("Collection name"),
			),
		),
		s.deleteCollection,
	)
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
				mcp.Description("Source type tag (convenience alias for metadata.source_type)"),
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
	s.server.AddTool(
		mcp.NewTool(
			constant.Tag,
			mcp.WithDescription(
				"Set, get, or remove a source type tag for a path prefix. Pass source_type to set, empty string to remove, omit to get.",
			),
			mcp.WithString(
				constant.Path,
				mcp.Required(),
				mcp.Description("Path prefix to tag"),
			),
			mcp.WithString(
				constant.Collection,
				mcp.Description("Scope tag to a collection (omit for global)"),
			),
			mcp.WithString(
				constant.SourceType,
				mcp.Description("Source type to set (omit to get, empty string to remove)"),
			),
		),
		s.tag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTags,
			mcp.WithDescription(
				"List all configured source type tags.",
			),
		),
		s.listTags,
	)
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
				mcp.Description("Specific key to expand (shows all values regardless of cardinality)"),
			),
		),
		s.listMetadata,
	)
}
