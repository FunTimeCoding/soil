package gomemoryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	queryConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Memory,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := store.New(connection.New(l, o.LitePath))
	defer s.Close()
	q := client.NewEnvironment()
	defer q.Close()
	errors.PanicOnError(
		peer.Await(
			context.Background(),
			l,
			queryConstant.Identity.Name(),
			q.Address(),
			q.Reachable,
		),
	)
	v := service.New(s, q).WithHiddenTag(
		environment.Fallback(constant.HiddenTagEnvironment, ""),
	).WithReferenceRoot(
		environment.Fallback(constant.ReferenceRootEnvironment, ""),
	)
	index := reacher.New(r)

	for _, f := range reconcileMemories(v) {
		index.Handle(
			queryConstant.Identity.Name(),
			f,
			func(x *edge.Edge) { l.Plain("reconcile: %s", x) },
			func(g error) { r.CaptureException(g) },
		)
	}

	u := web.New(v)
	lifecycle.New(l).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, u, r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(u.Recovery(r)),
	).RunUntilSignal()
}
