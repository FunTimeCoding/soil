package gofirefoxd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/firefox"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Firefox,
	r face.Reporter,
) {
	c := firefox.NewEnvironment()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				web.AddressPort(o.BridgePort),
				func(m *http.ServeMux) {
					m.Handle("/", c)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
		lifecycle.WithServer(
			server.New(
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(
						c,
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
