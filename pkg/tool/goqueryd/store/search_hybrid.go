package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"sort"
)

func (s *Store) SearchHybrid(
	o *search_option.Option,
	l *ollama.Client,
) ([]SearchResult, error) {
	scores := map[string]float64{}
	byPath := map[string]SearchResult{}
	bodyByPath := map[string]string{}
	addList := func(results []SearchResult) {
		for rank, r := range results {
			scores[r.FilePath] += 1.0 / float64(constant.RrfK+rank+1)
			existing, found := byPath[r.FilePath]

			if !found {
				byPath[r.FilePath] = r
			} else if r.ChunkPosition > 0 && existing.ChunkPosition == 0 {
				r.Score = existing.Score
				byPath[r.FilePath] = r
			}
		}
	}
	addBody := func(results []SearchResult) {
		for _, r := range results {
			if _, okay := bodyByPath[r.FilePath]; !okay {
				if r.Body != "" {
					bodyByPath[r.FilePath] = r.Body
				}
			}
		}
	}
	fetchFull := o.Full || o.Reranker != nil
	keywordResults, e := s.SearchKeyword(
		o.Query,
		o.Limit*2,
		o.Collection,
		fetchFull,
	)

	if e != nil {
		return nil, e
	}

	addList(keywordResults)
	addBody(keywordResults)
	vectorResults, f := s.SearchVector(
		o.Query,
		o.Limit*2,
		o.Collection,
		fetchFull,
		l,
	)

	if f != nil {
		return nil, f
	}

	addList(vectorResults)
	addBody(vectorResults)
	var expanded []ExpandedQuery

	if constant.ExpandModel != "" {
		var g error
		expanded, g = ExpandQuery(o.Query, l)

		if g != nil {
			return nil, g
		}
	}

	for _, q := range expanded {
		switch q.Type {
		case "lex":
			lexResults, h := s.SearchKeyword(
				q.Query,
				o.Limit*2,
				o.Collection,
				fetchFull,
			)

			if h != nil {
				return nil, h
			}

			addList(lexResults)
			addBody(lexResults)
		case "vec", "hyde":
			vecResults, h := s.SearchVector(
				q.Query,
				o.Limit*2,
				o.Collection,
				fetchFull,
				l,
			)

			if h != nil {
				return nil, h
			}

			addList(vecResults)
			addBody(vecResults)
		}
	}

	excludeSet := map[string]bool{}

	for _, p := range o.Exclude {
		excludeSet[p] = true
	}

	type scored struct {
		result SearchResult
		score  float64
	}
	merged := make([]scored, 0, len(scores))

	for path, score := range scores {
		r := byPath[path]

		if excludeSet[r.Path] {
			continue
		}

		r.Score = score
		r.Source = "hybrid"
		merged = append(merged, scored{result: r, score: score})
	}

	sort.Slice(
		merged,
		func(i, j int) bool {
			return merged[i].score > merged[j].score
		},
	)
	candidateLimit := o.Limit * 3

	if candidateLimit > len(merged) {
		candidateLimit = len(merged)
	}

	candidates := merged[:candidateLimit]

	if o.Reranker != nil && len(candidates) > 0 {
		documents := make([]string, len(candidates))

		for i, c := range candidates {
			body := bodyByPath[c.result.FilePath]

			if body == "" {
				body = c.result.Snippet
			}

			if len(body) > 2000 {
				body = body[:2000]
			}

			documents[i] = body
		}

		ranked, g := o.Reranker.Rank(o.Query, documents)

		if g == nil {
			reranked := make([]scored, len(candidates))

			for i, r := range ranked {
				reranked[i] = scored{
					result: candidates[r.Index].result,
					score:  r.Score,
				}
				reranked[i].result.Score = r.Score
				reranked[i].result.Source = "rerank"
			}

			sort.Slice(
				reranked,
				func(i, j int) bool {
					return reranked[i].score > reranked[j].score
				},
			)
			candidates = reranked
		}
	}

	unenriched := make([]SearchResult, 0, len(candidates))

	for _, m := range candidates {
		if !o.Full {
			m.result.Body = ""
		}

		unenriched = append(unenriched, m.result)
	}

	enriched := s.EnrichResults(unenriched, o.Metadata)

	if len(enriched) > o.Limit {
		enriched = enriched[:o.Limit]
	}

	return enriched, nil
}
