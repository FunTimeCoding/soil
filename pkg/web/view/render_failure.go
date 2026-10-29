package view

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) renderFailure(
	w http.ResponseWriter,
	q *http.Request,
	status int,
	item gomponents.Node,
) {
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)

	if v.IsExtendedRequest(q) {
		w.Header().Set(constant.NotificationItem, "true")
		w.WriteHeader(status)
		errors.LogOnError(item.Render(w))

		return
	}

	w.WriteHeader(status)
	page := v.layout.Clone().
		WithTitle("Error").
		WithPath(q.URL.Path).
		WithContent(item).
		Render()
	errors.LogOnError(page.Render(w))
}
