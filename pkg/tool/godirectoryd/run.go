package godirectoryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Directory,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	d := daemon.New(o, l, r)
	defer d.Close()
	errors.PanicOnError(d.Wait(context.Background()))
	c := lifecycle.New(l)

	for _, w := range d.Workers() {
		c.WithWorker(w)
	}

	c.WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				d.Mount(r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
