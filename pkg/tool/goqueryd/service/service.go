package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

type Service struct {
	store    *store.Store
	embedder face.Embedder
	reranker *rerank.Reranker
}
