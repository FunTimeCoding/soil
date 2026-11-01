package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
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
		entry_page.AddForm(url.Values{}, ""),
	)
}
