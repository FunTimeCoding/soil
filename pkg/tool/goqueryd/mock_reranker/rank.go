package mock_reranker

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"

func (r *Reranker) Rank(
	_ string,
	documents []string,
) ([]rerank.Result, error) {
	result := make([]rerank.Result, len(documents))

	for i := range documents {
		result[i] = rerank.Result{Index: i}
	}

	return result, nil
}
