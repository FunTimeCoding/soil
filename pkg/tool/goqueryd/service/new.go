package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func New(
	s *store.Store,
	m face.Embedder,
	re *rerank.Reranker,
) *Service {
	return &Service{store: s, embedder: m, reranker: re}
}
