package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Query struct {
	option.Web
	Metric       option.Web
	LitePath     string
	EmbedTarget  string
	EmbedModel   string
	Embedder     func() face.ReachableEmbedder
	Reranker     func() queryd.Reranker
	EmbedCounter func() queryd.EmbedCounter
}
