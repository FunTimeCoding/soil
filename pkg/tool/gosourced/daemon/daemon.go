package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"sync"
)

type Daemon struct {
	memory    face.Client
	guard     *peer.Guard
	service   *service.Service
	closeOnce sync.Once
}
