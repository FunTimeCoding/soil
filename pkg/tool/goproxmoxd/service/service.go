package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"sync"
)

type Service struct {
	inventory   *instance.Inventory
	pool        *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	shells      *pool.Pool[instance.Instance, *instance.Instance, ssh.SSHClient]
	netbox      netbox.Client
	netboxGuard *peer.Guard
	logger      *logger.Logger
	reporter    soilFace.Reporter
	notifier    *notifier.Notifier
	mutex       sync.RWMutex
	floor       *floor.Floor
}
