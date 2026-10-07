package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func fuse(
	exclude []string,
	lists ...[]result.Search,
) ([]result.Ranked, map[string]string) {
	scores := map[string]float64{}
	byPath := map[string]result.Search{}
	bodies := map[string]string{}

	for _, results := range lists {
		for rank, r := range results {
			scores[r.FilePath] += 1.0 / float64(constant.RrfK+rank+1)
			existing, found := byPath[r.FilePath]

			if !found {
				byPath[r.FilePath] = r
			} else if r.ChunkPosition >= 0 && existing.ChunkPosition < 0 {
				r.Score = existing.Score
				byPath[r.FilePath] = r
			}

			if _, okay := bodies[r.FilePath]; !okay && r.Body != "" {
				bodies[r.FilePath] = r.Body
			}
		}
	}

	excluded := map[string]bool{}

	for _, p := range exclude {
		excluded[p] = true
	}

	s := make([]result.Ranked, 0, len(scores))

	for path, score := range scores {
		r := byPath[path]

		if excluded[r.Path] {
			continue
		}

		r.Score = score
		r.Source = "hybrid"
		s = append(s, result.Ranked{Search: r, Score: score})
	}

	sortByScore(s)

	return s, bodies
}
