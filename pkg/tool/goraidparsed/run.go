package goraidparsed

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Parser,
	s face.Instrument,
) {
	r := s.Reporter()
	l := logger.New(context.Background())
	lifecycle.New(
		l,
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						server.New(
							o.ParserPath,
							o.TemplatePath,
							o.OutputPath,
							l,
							r,
						),
						s.Recorder(),
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
