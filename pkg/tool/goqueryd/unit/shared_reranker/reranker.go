package shared_reranker

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"

type Reranker struct {
	reranker *rerank.Reranker
}
