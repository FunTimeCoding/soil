package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func rerankCandidates(
	o *search_option.Option,
	candidates []search.Ranked,
	bodies map[string]string,
) []search.Ranked {
	if o.Reranker == nil || len(candidates) == 0 {
		return candidates
	}

	documents := make([]string, len(candidates))

	for i, c := range candidates {
		documents[i] = RerankText(
			bodies[c.Result.FilePath],
			&c.Result,
			o.Reranker,
		)
	}

	ranked, e := o.Reranker.Rank(o.Query, documents)

	if e != nil {
		return candidates
	}

	result := make([]search.Ranked, len(candidates))

	for i, r := range ranked {
		result[i] = search.Ranked{
			Result: candidates[r.Index].Result,
			Score:  r.Score,
		}
		result[i].Result.Score = r.Score
		result[i].Result.Source = "rerank"
	}

	sortByScore(result)

	return result
}
