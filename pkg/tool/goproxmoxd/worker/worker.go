package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"sync"
	"time"
)

type Worker struct {
	client    func(string) (face.ProxmoxClient, error)
	inventory *instance.Inventory
	interval  time.Duration
	collector *collector.Collector
	log       *logger.Logger
	recovery  *recovery.Recovery
	notifier  *notifier.Notifier
	reacher   *reacher.Reacher
	stop      chan struct{}
	done      chan struct{}
	started   sync.Once
	stopped   sync.Once
	mutex     sync.RWMutex
	floor     *floor.Floor
}
