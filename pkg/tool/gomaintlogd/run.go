package gomaintlogd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Log,
	i face.Instrument,
) {
	r := i.Reporter()
	g := logger.New(context.Background())
	s := service.New(o, g)
	defer s.Close()
	v := web.New(s, r)
	lifecycle.New(g).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(s, v, r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(v.Recovery(r)),
	).RunUntilSignal()
}
