package daemon

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/worker"
	"sync"
)

type Daemon struct {
	database  *sql.DB
	query     face.Client
	guard     *peer.Guard
	service   *service.Service
	web       *web.Server
	clock     *worker.Worker
	hiddenTag string
	closeOnce sync.Once
}
