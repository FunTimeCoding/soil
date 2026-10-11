package base

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
)

func newReranker() *rerank.Reranker {
	a, e := rerank.NewEnvironment()
	errors.PanicOnError(e)

	return a
}
