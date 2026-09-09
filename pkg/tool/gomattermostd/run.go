package gomattermostd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	mattermostFace "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	monitorOption "github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/sweeper"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"github.com/funtimecoding/soil/pkg/web"
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
	var d *store.Store
	var index mattermostFace.Indexer
	var p []lifecycle.Option

	if claude := connector.NewOptional(); claude != nil {
		d = store.New(relational.Open(l, o.PostgresLocator, o.LitePath))
		defer d.Close()
		n := notifier.New(claude, constant.Identity.Name(), r)
		w := watcher.New(c, d, n, l, r, constant.DebounceWindow)
		index = w
		p = append(
			p,
			lifecycle.WithWorker(w),
			lifecycle.WithWorker(
				sweeper.New(
					d,
					n,
					w,
					l,
					r,
					constant.PurgeWindow,
					constant.PurgeInterval,
				),
			),
		)
	}

	var m *monitor.Monitor

	if v := monitorOption.NewEnvironment(); v.Enabled {
		m = monitor.New(c, v, l, r)
		p = append(p, lifecycle.WithWorker(m))
	}

	lifecycle.New(
		l,
		append(
			p,
			lifecycle.WithServer(
				server.New(
					constant.Identity,
					o.Address,
					func(u *http.ServeMux) {
						Mount(
							c,
							m,
							d,
							index,
							r,
							s.Recorder(),
							o.Version,
							guard.New(u, o.ServiceTokens),
						)
					},
				).WithMiddleware(web.RecoveryMiddleware(r)),
			),
		)...,
	).RunUntilSignal()
}
