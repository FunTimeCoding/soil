package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

type Query struct {
	Address       string
	MetricAddress string
	ServiceTokens []string
	LitePath      string
	EmbedTarget   string
	EmbedModel    string
	Embedder      func() face.ReachableEmbedder
	Reranker      func() queryd.Reranker
	EmbedCounter  func() queryd.EmbedCounter
}
