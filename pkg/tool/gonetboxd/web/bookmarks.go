package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) bookmarks(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result, e := s.client.Bookmarks()
	errors.PanicOnError(e)
	var content gomponents.Node

	if len(result) == 0 {
		content = html.P(gomponents.Text(constant.BookmarkNone))
	} else {
		content = bookmarkTable(result)
	}

	s.view.RenderPageWithSummary(
		w,
		constant.BookmarkTitle,
		constant.BookmarkPath,
		summary(result),
		content,
	)
}
