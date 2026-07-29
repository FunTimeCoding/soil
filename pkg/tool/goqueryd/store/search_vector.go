package store

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"sort"
)

func (s *Store) SearchVector(
	query string,
	limit int,
	collection string,
	full bool,
	metadata map[string]string,
	o *ollama.Client,
) ([]SearchResult, error) {
	queryVector, e := o.EmbedSingle(constant.OllamaEmbedModel, query)

	if e != nil {
		return nil, e
	}

	candidates := s.allEmbeddings(collection, withoutSourceType(metadata))

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
			Context:       s.ResolveContext(c.collection, c.path),
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
