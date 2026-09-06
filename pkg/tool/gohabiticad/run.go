package gohabiticad

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/habitica"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Habitica,
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
					Mount(
						habitica.NewEnvironment(),
						r,
						s.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
