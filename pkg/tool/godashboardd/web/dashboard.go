package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	values := s.service.Values()
	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		s.topGrid(values),
		s.tailGrid(values),
	)
}
