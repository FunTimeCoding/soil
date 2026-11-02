package daemon

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

type Daemon struct {
	inventory *instance.Inventory
	pool      *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	selection *selection.Selection
	service   *service.Service
}
