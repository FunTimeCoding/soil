package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func rerankCandidates(
	o *search_option.Option,
	candidates []result.Ranked,
	bodies map[string]string,
) []result.Ranked {
	if o.Reranker == nil || len(candidates) == 0 {
		return candidates
	}

	documents := make([]string, len(candidates))

	for i, c := range candidates {
		documents[i] = RerankText(
			bodies[c.Search.FilePath],
			&c.Search,
			o.Reranker,
		)
	}

	ranked, e := o.Reranker.Rank(o.Query, documents)

	if e != nil {
		return candidates
	}

	s := make([]result.Ranked, len(candidates))

	for i, r := range ranked {
		s[i] = result.Ranked{
			Search: candidates[r.Index].Search,
			Score:  r.Score,
		}
		s[i].Search.Score = r.Score
		s[i].Search.Source = "rerank"
	}

	sortByScore(s)

	return s
}
