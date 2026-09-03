package goterraformd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Terraform,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := store.New(
		relational.Open(l, o.PostgresLocator, o.LitePath),
		"terraform_runs",
	)
	p := reaper.New(r)
	m := metric.New()
	n := runner.New(
		o,
		s,
		l,
		r,
		p,
		m.Registry(),
		client.NewInCluster(kubernetesConstant.InCluster),
	)
	lifecycle.New(
		l,
		lifecycle.WithWorker(p),
		lifecycle.WithWorker(n),
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
					model_context.New(n, s, r, i.Recorder(), o.Version).Mount(
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
