package web

import (
	tool "github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(constant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(constant.RootPattern), s.board)
	g.Open(route.Get(tool.PipelinePath), s.pipeline)
	g.Open(route.Get(tool.JobPath), s.job)
	g.Open(route.Post(tool.RetryPath), s.retry)
	g.Open(route.Post(tool.DeletePath), s.delete)
	g.OpenMount(route.Get(constant.LivePath), s.event())
	g.OpenMount(route.Get("/static/"), http.FileServerFS(staticFiles))
	g.Open(route.Get(constant.FaviconPath), s.favicon)
}
