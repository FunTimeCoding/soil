package gocertificated

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Certificate,
	u face.Instrument,
) {
	r := u.Reporter()
	g := logger.New(context.Background())
	v := service.New(o, g, r)
	defer v.Close()
	i := web.New(v, r)
	lifecycle.New(g).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, i, r, u.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(i.Recovery(r)),
	).RunUntilSignal()
}
