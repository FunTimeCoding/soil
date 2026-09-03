package goitermd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/iterm"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Iterm,
	s face.Instrument,
) {
	r := s.Reporter()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(
						iterm.NewEnvironment(),
						r,
						s.Recorder(),
						o.Version,
					).Mount(guard.New(m, o.ServiceTokens))
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
