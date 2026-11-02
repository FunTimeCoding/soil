package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
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
	p := pool.New(o.Inventory, peeredBuilder(o, l, r))

	for _, i := range o.Inventory.Instances {
		_, e := p.Client(i.Name)
		errors.PanicOnError(e)
	}

	shells := pool.New(o.Inventory, o.Shell)
	v := service.New(p, shells, o.Inventory, l, r)
	result := &Daemon{
		inventory: o.Inventory,
		pool:      p,
		shells:    shells,
		selection: selection.New(),
		service:   v,
		worker:    worker.New(v.Refresh, constant.PollInterval, y, l, r),
		web:       web.New(o.Inventory, v),
	}

	if o.Netbox == nil {
		return result
	}

	result.netbox = o.Netbox()
	result.netboxGuard = peer.New(
		gonetboxdConstant.Identity.Name(),
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
