package mock_reranker

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_result"

func (r *Reranker) Rank(
	_ string,
	documents []string,
) ([]*rerank_result.Result, error) {
	result := make([]*rerank_result.Result, len(documents))

	for i := range documents {
		result[i] = rerank_result.New(i, float64(len(documents)-i))
	}

	return result, nil
}
