package shared_reranker

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_result"

func (r *Reranker) Rank(
	query string,
	documents []string,
) ([]*rerank_result.Result, error) {
	return r.reranker.Rank(query, documents)
}
