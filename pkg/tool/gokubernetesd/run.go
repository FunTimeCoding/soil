package gokubernetesd

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
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
	svc := service.New(store.New(o.Path))

	if !svc.HasClusters() {
		fmt.Fprintln(os.Stderr, "no kubernetes clusters available")
		os.Exit(1)
	}

	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				web.AddressPort(o.Port),
				func(m *http.ServeMux) {
					model_context.New(
						svc,
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
