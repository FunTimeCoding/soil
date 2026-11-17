package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
)

func (o *Tester) daemonOption() *option.Query {
	result := option.New()
	result.LitePath = o.Path
	result.EmbedTarget = constant.FixtureEmbedTarget
	result.EmbedModel = o.EmbedModel
	result.Embedder = func() face.ReachableEmbedder { return o.embedder }
	result.Reranker = func() queryd.Reranker { return o.reranker }
	result.EmbedCounter = func() queryd.EmbedCounter { return o.counter }

	return result
}
