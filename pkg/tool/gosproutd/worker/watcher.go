package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/types/discovered_file"
	"sync"
)

type Watcher struct {
	sync          func([]discovered_file.File)
	recovery      *recovery.Recovery
	reporter      face.Reporter
	seedDirectory string
	stop          chan struct{}
	closeOnce     sync.Once
	running       sync.WaitGroup
}
