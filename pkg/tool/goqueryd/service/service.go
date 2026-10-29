package service

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"sync"
)

type Service struct {
	database  *sql.DB
	store     *store.Store
	embedder  face.ReachableEmbedder
	guard     *peer.Guard
	reranker  queryd.Reranker
	closeOnce sync.Once
}
