package goatlassiand

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Atlassian,
	s face.Instrument,
) {
	r := s.Reporter()
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
				d.Mount(r, s.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
