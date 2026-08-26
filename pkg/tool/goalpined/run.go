package goalpined

import (
	"context"
	alpine "github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Alpine,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := package_server.NewEnvironment()
	c := model_context.New(r, i.Recorder(), o.Version)
	lifecycle.New(
		l,
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					s.Mount(m)
					c.Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				constant.FileAddress,
				func(m *http.ServeMux) {
					m.Handle(
						strings.Slash,
						http.FileServer(http.Dir(alpine.PackageRoot)),
					)
				},
			),
		),
	).RunUntilSignal()
}
