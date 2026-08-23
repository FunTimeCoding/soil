package gogitlabd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Gitlab,
	s face.Instrument,
) {
	r := s.Reporter()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(
						gitlab.NewEnvironment().Nested(),
						r,
						s.Recorder(),
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
