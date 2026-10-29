package view

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderPageWithStatus(
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
	errors.PanicOnError(page.Render(w))
}
