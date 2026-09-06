package goatlassiand

import (
	"github.com/funtimecoding/soil/pkg/face"
	atlassianFace "github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
	soilWeb "github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	j atlassianFace.JiraSource,
	c atlassianFace.ConfluenceSource,
	b *web.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(j, c, r),
				[]generated.StrictMiddlewareFunc{
					soilWeb.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(j, c, r, t, version).Mount(g)
	b.Mount(g)
}
