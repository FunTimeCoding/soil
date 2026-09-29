package golinkaced

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/linkace"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/option"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.LinkAce,
	c *linkace.Client,
	i face.Instrument,
) {
	r := i.Reporter()
	s := service.New(c)
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						c,
						s,
						r,
						i.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
