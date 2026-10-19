package watcher

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
)

type Watcher struct {
	service       *service.Service
	recovery      *recovery.Recovery
	reporter      face.Reporter
	seedDirectory string
	stop          chan struct{}
}
