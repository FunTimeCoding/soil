package gokubernetesd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"os"
)

func Run(
	o *option.Server,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := service.New(o, l, r)

	if !s.HasClusters() {
		s.Close()
		errors.Println("no kubernetes clusters available")
		os.Exit(1)
	}

	defer s.Close()
	c := lifecycle.New(l)

	for _, w := range s.Workers() {
		c.WithWorker(w)
	}

	c.WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(
					s,
					o.ReadOnly,
					r,
					i.Recorder(),
					guard.New(m, o.ServiceTokens),
				)
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).RunUntilSignal()
}
