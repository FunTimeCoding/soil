package daemon

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/worker"
	"sync"
)

type Daemon struct {
	database  *sql.DB
	embedder  face.ReachableEmbedder
	guard     *peer.Guard
	reranker  queryd.Reranker
	counter   queryd.EmbedCounter
	service   *service.Service
	web       *web.Server
	clock     *worker.Worker
	closeOnce sync.Once
}
