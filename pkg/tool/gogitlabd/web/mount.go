package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), s.registry.Serve())
	g.Open(route.Get(webConstant.RootPattern), s.board)
	g.Open(route.Get(constant.PipelinePath), s.pipeline)
	g.Open(route.Get(constant.JobPath), s.job)
	g.Open(route.Post(constant.RetryPath), s.retry)
	g.Open(route.Post(constant.DeletePath), s.delete)
	g.OpenMount(route.Get(webConstant.LivePath), s.event())
	g.OpenMount(route.Get("/static/"), http.FileServerFS(staticFiles))
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
