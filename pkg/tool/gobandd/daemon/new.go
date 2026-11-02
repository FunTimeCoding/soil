package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func New(
	o *option.Band,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	p := pool.New(o.Inventory, peeredBuilder(o, l, r))

	for _, i := range o.Inventory.Instances {
		_, e := p.Client(i.Name)
		errors.PanicOnError(e)
	}

	result := &Daemon{
		inventory: o.Inventory,
		pool:      p,
		selection: selection.New(),
		service:   service.New(p, l, r),
	}

	if o.Netbox != nil {
		result.netbox = o.Netbox()
		result.netboxGuard = peer.New(
			gonetboxdConstant.Identity.Name(),
			result.netbox.Address(),
			result.netbox,
			l,
			r,
		)
		result.service.WithNetbox(result.netbox, result.netboxGuard)
	}

	return result
}
