package service

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/embed_model"
	"sync"
)

type Service struct {
	model     *embed_model.Model
	database  *sql.DB
	store     *store.Store
	embedder  face.ReachableEmbedder
	guard     *peer.Guard
	reranker  queryd.Reranker
	counter   queryd.EmbedCounter
	closeOnce sync.Once
}
