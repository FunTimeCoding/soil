package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	netboxConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/reconciler"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/web"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Proxmox,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	p := pool.New(o.Inventory, o.Hypervisor)

	for _, i := range o.Inventory.Instances {
		_, e := p.Client(i.Name)
		errors.PanicOnError(e)
	}

	shells := pool.New(o.Inventory, o.Shell)
	k := worker.New(p.Client, o.Inventory, constant.PollInterval, y, l, r)
	result := &Daemon{
		inventory: o.Inventory,
		pool:      p,
		shells:    shells,
		selection: selection.New(),
		service:   service.New(p, shells, o.Inventory, l, r),
		worker:    k,
		web:       web.New(o.Inventory, k),
	}

	if o.Netbox == nil {
		return result
	}

	result.netbox = o.Netbox()
	result.netboxGuard = peer.New(
		netboxConstant.Identity.Name(),
		result.netbox.Address(),
		result.netbox,
		l,
		r,
	)
	result.service.WithNetbox(result.netbox, result.netboxGuard)
	result.reconciler = reconciler.New(
		result.service,
		o.Inventory,
		constant.ReconcileInterval,
		l,
		r,
	)

	return result
}
