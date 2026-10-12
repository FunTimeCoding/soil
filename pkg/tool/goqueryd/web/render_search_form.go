package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) renderSearchForm(
	w http.ResponseWriter,
	collection string,
) {
	status := s.service.MustStatus()
	s.view.RenderPage(
		w,
		constant.SearchTitle,
		webConstant.SearchPath,
		html.H3(gomponents.Text(constant.SearchTitle)),
		searchForm("", collection, status.Collections),
	)
}
