package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/document_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) documentDetailPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	path := q.PathValue(constant.Path)
	document, suggestions, e := s.service.GetDocument(path)

	if not_found.Is(e) {
		s.view.RenderPageWithStatus(
			w,
			http.StatusNotFound,
			"Not Found",
			webConstant.SearchPath,
			document_page.Missing(path, suggestions),
		)

		return
	}

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		document.Title,
		constant.CollectionsPath,
		document_page.Detail(document),
	)
}
