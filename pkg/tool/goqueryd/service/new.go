package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func New(
	s *store.Store,
	m face.ReachableEmbedder,
	re queryd.Reranker,
	target string,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	return &Service{
		store:    s,
		embedder: m,
		guard:    peer.New(constant.EmbedderPeer, target, m, l, r),
		reranker: re,
	}
}
