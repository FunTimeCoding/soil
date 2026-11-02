package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Gitlab,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	p := pool.New(o.Inventory, peeredBuilder(o, l, r))

	for _, i := range o.Inventory.Instances {
		_, e := p.Client(i.Name)
		errors.PanicOnError(e)
	}

	s := service.New(p, o.Inventory)

	return &Daemon{
		inventory: o.Inventory,
		pool:      p,
		selection: selection.New(),
		service:   s,
		web:       web.New(s, o.Inventory, r),
		worker: worker.New(
			s.Refresh,
			s.Entries,
			constant.PollInterval,
			y,
			l,
			r,
		),
	}
}
