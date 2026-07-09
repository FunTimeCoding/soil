package gokubernetesd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"os"
)

func Run(
	o *option.Server,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := service.New(store.New(lite.New(l, o.LitePath)))

	if !s.HasClusters() {
		errors.Println("no kubernetes clusters available")
		os.Exit(1)
	}

	lifecycle.New(
		l,
		lifecycle.WithServer(
			server.New(
				o.Address,
				func(m *http.ServeMux) {
					model_context.New(
						s,
						o.ReadOnly,
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
