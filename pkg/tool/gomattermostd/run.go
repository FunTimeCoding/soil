package gomattermostd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/server"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Mattermost,
	s face.Instrument,
) {
	r := s.Reporter()
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
				lifecycleServer.New(
					constant.Identity,
					o.Address,
					func(u *http.ServeMux) {
						guard.New(u, o.ServiceTokens).TokenMount(
							webConstant.InterfacePath,
							generated.HandlerFromMux(
								generated.NewStrictHandler(
									server.New(c, o.Version, r),
									nil,
								),
								http.NewServeMux(),
							),
						)
						model_context.New(
							c,
							m,
							r,
							s.Recorder(),
							o.Version,
						).Mount(guard.New(u, o.ServiceTokens))
					},
				).WithMiddleware(web.RecoveryMiddleware(r)),
			),
		)...,
	).RunUntilSignal()
}
