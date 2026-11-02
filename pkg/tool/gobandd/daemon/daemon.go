package daemon

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Daemon struct {
	inventory   *instance.Inventory
	pool        *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	selection   *selection.Selection
	netbox      face.Client
	netboxGuard *peer.Guard
	service     *service.Service
}
