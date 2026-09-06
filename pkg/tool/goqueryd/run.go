package goqueryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"time"
)

func Run(
	o *option.Query,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := store.New(connection.New(l, o.LitePath))
	defer s.Close()
	a, e := rerank.New(o.RerankModel, o.RerankTokenizer)
	errors.PanicOnError(e)
	defer errors.LogClose(a)
	v := service.New(s, embed.NewEnvironment(), a)
	u := web.New(v)
	lifecycle.New(
		l,
		lifecycle.WithWorker(worker.New(v, 10*time.Minute, l, r)),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						v,
						u,
						r,
						i.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
