package store

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
	"sort"
)

func (s *Store) SearchVector(
	query string,
	limit int,
	collection string,
	full bool,
	metadata map[string]string,
	m face.Embedder,
) ([]search.Result, error) {
	queryVector, e := embed.Single(m, query)

	if e != nil {
		return nil, e
	}

	candidates := s.allEmbeddings(collection, withoutSourceType(metadata))

	for i := range candidates {
		candidates[i].Distance = cosineDistance(
			queryVector,
			candidates[i].Vector,
		)
	}

	sort.Slice(
		candidates,
		func(i, j int) bool {
			return candidates[i].Distance < candidates[j].Distance
		},
	)
	seen := map[string]bool{}
	var result []search.Result

	for _, c := range candidates {
		if seen[c.FilePath] {
			continue
		}

		seen[c.FilePath] = true
		r := search.Result{
			VirtualPath:   buildVirtualPath(c.Collection, c.Path),
			FilePath:      c.FilePath,
			Collection:    c.Collection,
			Path:          c.Path,
			Title:         c.Title,
			Hash:          c.Hash,
			Score:         1 - c.Distance,
			Source:        "vec",
			Context:       s.ResolveContext(c.Collection, c.Path),
			ChunkPosition: c.Position,
		}
		snippet, line := ExtractSnippet(c.Body, query, r.ChunkPosition)
		r.Snippet = snippet
		r.SnippetLine = line

		if full {
			r.Body = c.Body
		}

		result = append(result, r)

		if len(result) >= limit {
			break
		}
	}

	return result, nil
}
