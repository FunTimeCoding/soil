package base

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"sync"
)

var (
	rerankerOnce     sync.Once
	rerankerInstance *rerank.Reranker
)

// Shared across the binary's tests - loading costs ~500ms and
// inference is read-only. Never closed; process exit reaps it.
func sharedReranker() *rerank.Reranker {
	rerankerOnce.Do(
		func() {
			a, e := rerank.New(
				environment.Required(constant.ModelEnvironment),
				environment.Required(constant.TokenizerEnvironment),
			)
			errors.PanicOnError(e)
			rerankerInstance = a
		},
	)

	return rerankerInstance
}
