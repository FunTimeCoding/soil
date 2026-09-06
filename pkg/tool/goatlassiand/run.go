package goatlassiand

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Atlassian,
	s face.Instrument,
) {
	r := s.Reporter()
	j := jira.NewEnvironment()
	c := confluence.NewEnvironment()
	l := logger.New(context.Background())
	k := worker.New(j, c, constant.PollInterval, l, r)
	b := web.New(k)
	lifecycle.New(
		l,
		lifecycle.WithWorker(k),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						j,
						c,
						b,
						r,
						s.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(b.Recovery(r)),
		),
	).RunUntilSignal()
}
