package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
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
	m.HandleFunc(route.Get(webConstant.SearchPath), s.searchPage)
	m.HandleFunc(route.Get(constant.CollectionsPath), s.collectionsPage)
	m.HandleFunc(
		route.Get(constant.CollectionsPath, "/{name}"),
		s.documentsPage,
	)
	m.HandleFunc(route.Get("/documents/{path...}"), s.documentDetailPage)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
}
