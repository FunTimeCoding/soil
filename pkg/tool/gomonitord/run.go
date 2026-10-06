package gomonitord

import (
	"context"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/option"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Monitor,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	n := notifier.New()
	s := store.New(relational.Open(l, o.PostgresLocator, o.LitePath), n)
	defer s.Close()
	lifecycle.New(
		l,
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(s, n, r, i.Recorder(), guard.New(m, o.ServiceTokens))
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
