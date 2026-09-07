package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

type Service struct {
	store    *store.Store
	embedder face.Embedder
	reranker queryd.Reranker
}
