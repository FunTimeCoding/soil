package goagentd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Agent,
	i face.Instrument,
) {
	r := i.Reporter()
	n := runner.New(o.Workspace)
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(n, i.Recorder(), guard.New(m, o.ServiceTokens))
				},
			).WithMiddleware(web.RecoveryMiddleware(r)).WithProtected(),
		),
	).RunUntilSignal()
}
