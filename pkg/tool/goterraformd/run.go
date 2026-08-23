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
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func Run(
	o *option.Terraform,
	r face.Reporter,
) {
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
				webConstant.MetricAddress,
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
					model_context.New(
						n,
						s,
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
