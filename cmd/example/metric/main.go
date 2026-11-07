package main

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	metricConstant "github.com/funtimecoding/soil/pkg/metric/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/option"
)

func main() {
	m := metric.New()
	lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(
			identity.Example(),
			option.New(
				join.Empty(
					stringsConstant.Colon,
					integers.ToString(metricConstant.Port),
				),
			),
			memory.New(),
			func(g *guard.Mux) {
				g.OpenMount(webConstant.MetricsPath, m.Exporter())
			},
		),
	).RunUntilSignal()
}
