package face

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"

type Reranker interface {
	Rank(
		query string,
		documents []string,
	) ([]rerank.Result, error)
}
