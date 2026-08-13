package base

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

type Server struct {
	t        *testing.T
	store    *store.Store
	embedder face.Embedder
	reranker *rerank.Reranker
	server   *model_context_server.Server
}
