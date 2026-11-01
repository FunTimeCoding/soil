package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/web/bookmark_page"
	"net/http"
)

func (s *Server) bookmarks(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result, e := s.service.Bookmarks()
	errors.PanicOnError(e)
	s.view.RenderPageWithSummary(
		w,
		constant.BookmarkTitle,
		constant.BookmarkPath,
		bookmark_page.Summary(result),
		bookmark_page.Table(result),
	)
}
