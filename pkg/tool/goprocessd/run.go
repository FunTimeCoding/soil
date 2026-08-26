package goprocessd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"os"
)

func Run(
	o *option.Option,
	i face.Instrument,
) {
	entries, e := procfile.Parse(o.ProcfilePath)
	errors.PanicOnError(e)
	env := environment.New(os.Environ())

	if f := env.Load(o.EnvrcPath); f != nil {
		errors.Printf("warning: %s\n", f)
	}

	s := server.New(
		entries,
		env,
		o.ProcfilePath,
		o.EnvrcPath,
		socket.Path(o.ProcfilePath),
	)
	r := i.Reporter()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(s, r, i.Recorder(), o.Version).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	)
	l.Run()

	defer l.Stop()
	errors.PanicOnError(s.Run())
}
