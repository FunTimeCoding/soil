package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

type Service struct {
	store    *store.Store
	embedder face.ReachableEmbedder
	guard    *peer.Guard
	reranker queryd.Reranker
}
