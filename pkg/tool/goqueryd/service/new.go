package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func New(
	embedModel string,
	s *store.Store,
	m face.ReachableEmbedder,
	g *peer.Guard,
	a queryd.Reranker,
	c queryd.EmbedCounter,
) *Service {
	return &Service{
		model:    embedModelByName(embedModel),
		store:    s,
		embedder: m,
		guard:    g,
		reranker: a,
		counter:  c,
	}
}
