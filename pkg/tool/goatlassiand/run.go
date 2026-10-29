package goatlassiand

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Atlassian,
	s face.Instrument,
) {
	r := s.Reporter()
	l := logger.New(context.Background())
	v := service.New(o, l, r)
	defer v.Close()
	errors.PanicOnError(v.Wait(context.Background()))
	k := worker.New(v, constant.PollInterval, l, r)
	b := web.New(k)
	c := lifecycle.New(l)

	for _, w := range v.Workers() {
		c.WithWorker(w)
	}

	c.WithWorker(k).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, b, r, s.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(b.Recovery(r)),
	).RunUntilSignal()
}
