package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) entryMissing(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	s.view.RenderPage(
		w,
		constant.EntryTitle,
		constant.EntriesPath,
		html.H1(gomponents.Text(constant.EntryTitle)),
		html.P(html.Em(gomponents.Text("This entry no longer exists."))),
		html.A(
			gomponents.Attr("role", "button"),
			html.Class("outline secondary"),
			html.Href(constant.EntriesPath),
			gomponents.Text(constant.EntriesTitle),
		),
	)
}
