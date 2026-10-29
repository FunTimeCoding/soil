package gosproutd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Sprout,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	v := service.New(o, l)
	defer v.Close()
	w := worker.New(v, l, r, o.SeedDirectory)
	u := web.New(v)
	lifecycle.New(l).WithWorker(w).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, u, r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(u.Recovery(r)).WithDefaultCertificate(),
	).RunUntilSignal()
}
