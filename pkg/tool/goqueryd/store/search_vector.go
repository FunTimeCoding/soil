package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"sort"
)

func (s *Store) SearchVector(
	query string,
	limit int,
	collection string,
	full bool,
	o *ollama.Client,
) ([]SearchResult, error) {
	queryVector, e := o.EmbedSingle(constant.EmbedModel, query)

	if e != nil {
		return nil, e
	}

	candidates := s.allEmbeddings(collection)

	for i := range candidates {
		candidates[i].distance = cosineDistance(
			queryVector,
			candidates[i].vector,
		)
	}

	sort.Slice(
		candidates,
		func(i, j int) bool {
			return candidates[i].distance < candidates[j].distance
		},
	)
	seen := map[string]bool{}
	var result []SearchResult

	for _, c := range candidates {
		if seen[c.filePath] {
			continue
		}

		seen[c.filePath] = true
		r := SearchResult{
			VirtualPath:   buildVirtualPath(c.collection, c.path),
			FilePath:      c.filePath,
			Collection:    c.collection,
			Path:          c.path,
			Title:         c.title,
			Hash:          c.hash,
			Score:         1 - c.distance,
			Source:        "vec",
			Context:       s.resolveContext(c.collection, c.path),
			ChunkPosition: c.position,
		}
		snippet, line := ExtractSnippet(c.body, query, r.ChunkPosition)
		r.Snippet = snippet
		r.SnippetLine = line

		if full {
			r.Body = c.body
		}

		result = append(result, r)

		if len(result) >= limit {
			break
		}
	}

	return result, nil
}
