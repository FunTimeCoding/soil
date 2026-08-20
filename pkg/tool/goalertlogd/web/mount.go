package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc(
		route.Get(webConstant.PalettePath),
		palette.NewServe(s.registry),
	)
	m.HandleFunc(route.Get(webConstant.RootPattern), s.dashboard)
	m.HandleFunc(route.Get(constant.RecentPath), s.recent)
	m.HandleFunc(route.Get("/alerts"), s.alerts)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
}
