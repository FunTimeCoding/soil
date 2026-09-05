package base

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
)

func newReranker() *rerank.Reranker {
	a, e := rerank.New(
		environment.Required(constant.ModelEnvironment),
		environment.Required(constant.TokenizerEnvironment),
	)
	errors.PanicOnError(e)

	return a
}
