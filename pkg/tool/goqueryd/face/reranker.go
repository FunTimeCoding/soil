package face

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_result"
)

type Reranker interface {
	face.TokenCounter
	Rank(
		query string,
		documents []string,
	) ([]*rerank_result.Result, error)
	Name() string
	SequenceLength() int
}
