package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc(route.Get(constant.PalettePath), palette.NewServe(s.registry))
	m.HandleFunc(route.Get(constant.RootPattern), s.plate)
	m.Handle(route.Get(constant.LivePath), s.event())
	m.HandleFunc(route.Get(constant.FaviconPath), s.favicon)
}
