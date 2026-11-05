package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/web/bookmark_page"
	"net/http"
)

func (s *Server) bookmarks(
	w http.ResponseWriter,
	r *http.Request,
) {
	result, e := s.service.Bookmarks()

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderPageWithSummary(
		w,
		constant.BookmarkTitle,
		constant.BookmarkPath,
		bookmark_page.Summary(result),
		bookmark_page.Table(result),
	)
}
