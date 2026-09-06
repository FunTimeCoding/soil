package goproxmoxd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/web"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Proxmox,
	s face.Instrument,
) {
	r := s.Reporter()
	v := service.New(o.Inventory)
	l := logger.New(context.Background())
	m := metric.New()
	k := worker.New(v, constant.PollInterval, m.Registry(), l, r)
	b := web.New(v, k)
	lifecycle.New(
		l,
		lifecycle.WithWorker(k),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.MetricAddress,
				func(x *http.ServeMux) {
					x.Handle(webConstant.MetricsPath, m.Exporter())
				},
			),
		),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						v,
						b,
						r,
						s.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(b.Recovery(r)),
		),
	).RunUntilSignal()
}
