package daemon

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
)

type Daemon struct {
	inventory *instance.Inventory
	pool      *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	selection *selection.Selection
	service   *service.Service
}
