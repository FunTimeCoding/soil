package gogithubd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func Run(
	o *option.Exporter,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	m := metric.New()
	lifecycle.New(
		l,
		lifecycle.WithVerbose(o.Verbose),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				webConstant.MetricAddress,
				func(x *http.ServeMux) {
					x.Handle(webConstant.MetricsPath, m.Exporter())
				},
			),
		),
		lifecycle.WithWorker(
			worker.New(
				github.NewEnvironment(),
				o.Owner,
				constant.PollInterval,
				m.Registry(),
				l,
				r,
			),
		),
	).RunUntilSignal()
}
