package view

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderStatusPage(
	w http.ResponseWriter,
	status int,
	title string,
	path string,
	content ...gomponents.Node,
) {
	page := v.layout.Clone().
		WithTitle(title).
		WithPath(path).
		WithContent(content...).
		Render()
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)
	w.WriteHeader(status)
	web.PanicOnWriteError(page.Render(w))
}
