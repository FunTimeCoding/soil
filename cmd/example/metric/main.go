package main

import (
	"context"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func main() {
	m := metric.New()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				identity.Example(),
				constant.MetricAddress,
				func(x *http.ServeMux) {
					x.Handle(constant.MetricsPath, m.Exporter())
				},
			),
		),
	).RunUntilSignal()
}
