package gofirefoxd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/service"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Firefox,
	s face.Instrument,
) {
	r := s.Reporter()
	v := service.New(o)
	defer v.Close()
	lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(
			constant.Identity,
			web.AddressPort(o.BridgePort),
			func(m *http.ServeMux) {
				m.Handle(webConstant.RootPath, v.Bridge())
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, r, s.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).RunUntilSignal()
}
