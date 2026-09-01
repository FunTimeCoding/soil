package gocredentiald

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/option"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"time"
)

func Run(
	o *option.Credential,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	v := service.New(o.Database, o.Password, o.RevealedField, time.Now, l)
	lifecycle.New(
		l,
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(v, r, i.Recorder(), o.Version).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
