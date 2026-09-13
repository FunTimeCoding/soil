package gomaintlogd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Log,
	i face.Instrument,
) {
	r := i.Reporter()
	g := logger.New(context.Background())
	events := notifier.New()
	s := store.New(relational.Open(g, o.PostgresLocator, o.LitePath), events)
	defer s.Close()
	v := web.New(s, events)
	lifecycle.New(
		g,
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						s,
						v,
						r,
						i.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(v.Recovery(r)),
		),
	).RunUntilSignal()
}
