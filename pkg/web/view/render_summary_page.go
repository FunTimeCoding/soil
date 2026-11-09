package view

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderSummaryPage(
	w http.ResponseWriter,
	title string,
	path string,
	summary []string,
	content ...gomponents.Node,
) {
	page := v.layout.Clone().
		WithTitle(title).
		WithPath(path).
		WithSummary(summary...).
		WithContent(content...).
		Render()
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)
	web.PanicOnWriteError(page.Render(w))
}
