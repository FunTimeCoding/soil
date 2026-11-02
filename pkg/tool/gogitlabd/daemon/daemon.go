package daemon

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
)

type Daemon struct {
	inventory *instance.Inventory
	pool      *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	selection *selection.Selection
	service   *service.Service
	web       *web.Server
	worker    *worker.Worker
}
