package gocertificated

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
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
	s := store.New(relational.Open(g, o.PostgresLocator, o.LitePath))
	defer s.Close()
	f := gitlab.NewEnvironment()
	project, e := f.ResolveProject(o.Project)
	errors.PanicOnError(e)
	v := service.New(
		s,
		publish.New(
			f,
			project,
			o.Branch,
			o.AuthorityDirectory,
			o.SecretAuthority,
			o.SecretPath,
		),
	)
	i := web.New(s, v, authorizationClient(o))
	lifecycle.New(
		g,
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						s,
						v,
						i,
						r,
						u.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(i.Recovery(r)),
		),
	).RunUntilSignal()
}
