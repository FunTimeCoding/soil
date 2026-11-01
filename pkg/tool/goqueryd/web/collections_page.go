package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/collection_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) collectionsPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	status, e := s.service.Status()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		constant.CollectionsTitle,
		constant.CollectionsPath,
		html.H3(gomponents.Text(constant.CollectionsTitle)),
		collection_page.List(status.Collections),
	)
}
