package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_embedder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"testing"
)

type Tester struct {
	Embedder   *mock_embedder.Embedder
	Reranker   *mock_reranker.Reranker
	Counter    *mock_reranker.Reranker
	Reporter   *memory.Memory
	Path       string
	EmbedModel string
	embedder   face.ReachableEmbedder
	reranker   queryd.Reranker
	counter    queryd.EmbedCounter
	service    *service.Service
	daemon     *daemon.Daemon
	t          *testing.T
}
