package gogitlabd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func Run(
	o *option.Gitlab,
	s face.Instrument,
) {
	r := s.Reporter()
	l := logger.New(context.Background())
	c := gitlab.NewEnvironment()
	m := metric.New()
	lifecycle.New(
		l,
		lifecycle.WithWorker(
			worker.New(c, constant.PollInterval, m.Registry(), l, r),
		),
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
				func(x *http.ServeMux) {
					model_context.New(
						c.Nested(),
						r,
						s.Recorder(),
						o.Version,
					).Mount(x)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
