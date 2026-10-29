package goraidd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Raid,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	v := service.New(o, l, r)
	defer v.Close()
	u := web.New(v, r)
	c := lifecycle.New(l)

	for _, w := range v.Workers() {
		c.WithWorker(w)
	}

	c.WithWorker(worker.New(v, l, r)).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, u, r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(u.Recovery(r)),
	).RunUntilSignal()
}
