package option

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/embed_counter"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
)

func New() *Query {
	return &Query{
		Embedder: embed.NewEnvironment,
		Reranker: func() face.Reranker {
			result, e := rerank.NewEnvironment()
			errors.PanicOnError(e)

			return result
		},
		EmbedCounter: func() face.EmbedCounter {
			result, e := embed_counter.NewEnvironment()
			errors.PanicOnError(e)

			return result
		},
	}
}
