package view

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderLivePageWithSummary(
	w http.ResponseWriter,
	title string,
	path string,
	liveParams string,
	summary []string,
	content ...gomponents.Node,
) {
	page := v.layout.Clone().
		WithTitle(title).
		WithPath(path).
		WithLiveParams(liveParams).
		WithSummary(summary...).
		WithContent(content...).
		Render()
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)
	errors.PanicOnError(page.Render(w))
}
