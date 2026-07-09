package gosentryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Server,
	r face.Reporter,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(
						sentry.NewEnvironment(),
						o.Organization,
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
