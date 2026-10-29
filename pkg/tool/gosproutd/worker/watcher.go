package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"sync"
)

type Watcher struct {
	service       *service.Service
	recovery      *recovery.Recovery
	reporter      face.Reporter
	seedDirectory string
	stop          chan struct{}
	closeOnce     sync.Once
}
