package goraidd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/raid_parser"
	raidParserConstant "github.com/funtimecoding/soil/pkg/raid_parser/constant"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func Run(
	o *option.Raid,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := store.New(
		relational.Open(l, o.PostgresLocator, o.LitePath),
		o.LogCachePath,
		o.ElitePath,
		l,
		r,
	)
	p := raid_parser.New(
		locator.Environment(
			raidParserConstant.HostEnvironment,
			raidParserConstant.PortEnvironment,
			raidParserConstant.InsecureEnvironment,
		),
		environment.Required(raidParserConstant.TokenEnvironment),
	)
	u := web.New(s, o.ElitePath, o.OutputPath, p, authorizationClient(o))
	lifecycle.New(
		l,
		lifecycle.WithWorker(s),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						s,
						o.OutputPath,
						u,
						r,
						i.Recorder(),
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
