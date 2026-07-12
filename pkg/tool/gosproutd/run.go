package gosproutd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/watcher"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"net/http"
)

func Run(
	o *option.Sprout,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(lite.New(l, o.LitePath))
	defer s.Close()
	v := service.New(s, notifier.New())
	w := watcher.New(v, l, r, o.SeedDirectory)
	u := web.New(v)
	lifecycle.New(
		l,
		lifecycle.WithWorker(w),
		lifecycle.WithServer(
			server.New(
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(
						v,
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
					u.Mount(m)
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
