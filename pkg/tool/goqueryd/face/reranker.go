package face

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
)

type Reranker interface {
	face.TokenCounter
	Rank(
		query string,
		documents []string,
	) ([]rerank.Result, error)
	Name() string
	SequenceLength() int
}
