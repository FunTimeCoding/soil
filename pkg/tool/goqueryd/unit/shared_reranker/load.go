package shared_reranker

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
)

func load() *rerank.Reranker {
	result, e := rerank.NewEnvironment()
	errors.PanicOnError(e)
	loaded.Store(true)

	return result
}
