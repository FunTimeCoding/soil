package goprocessd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Option,
	i face.Instrument,
) {
	v := service.New(o)
	r := i.Reporter()
	t := i.Recorder()
	l := lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, r, t, guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).WithServer(
		server.New(
			constant.Identity,
			"",
			func(m *http.ServeMux) { MountLocal(v, r, t, m) },
		).WithListener(socket.Listen(socket.Path(o.ProcfilePath))).
			WithMiddleware(web.RecoveryMiddleware(r)),
	)
	l.Run()

	defer l.Close()
	errors.PanicOnError(v.Run())
}
