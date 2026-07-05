package goterraformd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/provision/store"
	"github.com/funtimecoding/go-library/pkg/relational"
	"github.com/funtimecoding/go-library/pkg/system/reaper"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Terraform,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(
		relational.New(o.PostgresLocator).Mapper(),
		"terraform_runs",
	)
	n := runner.New(o, s, l, r)
	lifecycle.New(
		l,
		lifecycle.WithWorker(reaper.New(r)),
		lifecycle.WithWorker(n),
		lifecycle.WithServer(
			server.New(
				web.AddressPort(o.Port),
				func(m *http.ServeMux) {
					model_context.New(
						n,
						s,
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
