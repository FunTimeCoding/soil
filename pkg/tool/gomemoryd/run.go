package gomemoryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/memory_indexer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/connect"
	queryConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
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
	p := stall_client.New().Build()
	defer p.Close()
	idx := memory_indexer.New(connect.Wait(l, p.Client()))
	v := service.New(s, idx, idx, idx).WithHiddenTag(
		environment.Fallback(constant.HiddenTagEnvironment, ""),
	).WithReferenceRoot(
		environment.Fallback(constant.ReferenceRootEnvironment, ""),
	)
	index := reacher.New()

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
