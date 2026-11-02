package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/embed_model"
)

type Service struct {
	model    *embed_model.Model
	store    *store.Store
	embedder face.ReachableEmbedder
	guard    *peer.Guard
	reranker queryd.Reranker
	counter  queryd.EmbedCounter
}
