package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) sessions(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.SessionsTitle,
		constant.SessionsPath,
		html.H3(gomponents.Text(constant.SessionsTitle)),
		html.Div(extended.StreamSwap(constant.SessionsEvent), s.sessionTable()),
	)
}
