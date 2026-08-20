package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
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
	m.HandleFunc(route.Get(webConstant.RootPattern), s.heatmap)
	m.HandleFunc(route.Get(constant.EventsPath), s.events)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
}
