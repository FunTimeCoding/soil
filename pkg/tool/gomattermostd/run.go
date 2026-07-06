package gomattermostd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Mattermost,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	c := mattermost.NewEnvironment()
	var m *monitor.Monitor
	var p []lifecycle.Option

	if v := monitor.LoadConfiguration(); v.Enabled {
		m = monitor.New(c, v, l, r)
		p = append(p, lifecycle.WithWorker(m))
	}

	lifecycle.New(
		l,
		append(
			p,
			lifecycle.WithServer(
				server.New(
					web.AddressPort(o.Port),
					func(u *http.ServeMux) {
						model_context.New(
							c,
							m,
							r,
							telemetry.NewEnvironment(),
							o.Version,
						).Mount(u)
					},
				).WithMiddleware(web.RecoveryMiddleware(r)),
			),
		)...,
	).RunUntilSignal()
}
