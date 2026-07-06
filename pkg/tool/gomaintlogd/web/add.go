package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) add(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.AddEntryTitle,
		constant.AddEntryPath,
		html.H1(gomponents.Text(constant.AddEntryTitle)),
		html.Div(html.ID("result")),
		addForm(),
	)
}
