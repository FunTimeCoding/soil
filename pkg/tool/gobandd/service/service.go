package service

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Service struct {
	inventory   *instance.Inventory
	pool        *pool.Pool[instance.Instance, *instance.Instance, face.Engine]
	selection   *selection.Selection
	netbox      netbox.Client
	netboxGuard *peer.Guard
	logger      *logger.Logger
	reporter    soilFace.Reporter
}
