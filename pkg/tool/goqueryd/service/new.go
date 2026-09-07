package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func New(
	s *store.Store,
	m face.Embedder,
	re queryd.Reranker,
) *Service {
	return &Service{store: s, embedder: m, reranker: re}
}
