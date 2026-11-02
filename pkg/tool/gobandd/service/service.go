package service

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/peered"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Service struct {
	pool        *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	netbox      netbox.Client
	netboxGuard *peer.Guard
	logger      *logger.Logger
	reporter    soilFace.Reporter
}
