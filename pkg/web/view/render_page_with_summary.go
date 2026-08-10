package view

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderPageWithSummary(
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
	errors.PanicOnError(page.Render(w))
}
