package daemon

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/peer"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/reconciler"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/web"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
)

type Daemon struct {
	inventory   *instance.Inventory
	pool        *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	shells      *pool.Pool[instance.Instance, *instance.Instance, ssh.SSHClient]
	selection   *selection.Selection
	netbox      face.Client
	netboxGuard *peer.Guard
	service     *service.Service
	worker      *worker.Worker
	reconciler  *reconciler.Reconciler
	web         *web.Server
}
