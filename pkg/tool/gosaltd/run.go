package gosaltd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Salt,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(
		relational.Open(l, o.PostgresLocator, o.LitePath),
		"highstate_runs",
	)
	p := reaper.New(r)
	n := runner.New(o, salt.NewEnvironment, s, l, r, p)
	lifecycle.New(
		l,
		lifecycle.WithWorker(p),
		lifecycle.WithWorker(n),
		lifecycle.WithServer(
			server.New(
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
