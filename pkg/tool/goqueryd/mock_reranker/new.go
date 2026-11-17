package mock_reranker

import "math"

func New() *Reranker {
	return &Reranker{tokenLimit: math.MaxInt}
}
