package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func convertSearchResult(r store.SearchResult) *server.SearchResult {
	result := &server.SearchResult{
		Collection:  r.Collection,
		FilePath:    r.FilePath,
		Hash:        r.Hash,
		Path:        r.Path,
		Score:       float32(r.Score),
		Source:      r.Source,
		Title:       r.Title,
		VirtualPath: r.VirtualPath,
	}

	if r.Body != "" {
		result.Body = new(string(r.Body))
	}

	if r.Context != "" {
		result.Context = new(string(r.Context))
	}

	if r.Snippet != "" {
		result.Snippet = new(string(r.Snippet))
	}

	if r.SnippetLine != 0 {
		result.SnippetLine = new(int(r.SnippetLine))
	}

	if r.SourceType != "" {
		result.SourceType = new(string(r.SourceType))
	}

	if len(r.Metadata) > 0 {
		result.Metadata = &r.Metadata
	}

	return result
}
