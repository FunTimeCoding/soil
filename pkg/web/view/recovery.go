package view

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (v *View) Recovery(r face.Reporter) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				defer func() {
					caught := recover()

					if caught == nil {
						return
					}

					event := r.Recover(caught)
					item := layout.ErrorItem(
						fmt.Sprintf("%v", caught),
						event,
					)
					w.Header().Set(
						constant.ContentType,
						constant.MarkupUnicode,
					)

					if v.IsExtendedRequest(q) {
						w.Header().Set(constant.NotificationItem, "true")
						w.WriteHeader(http.StatusInternalServerError)
						errors.LogOnError(item.Render(w))

						return
					}

					w.WriteHeader(http.StatusInternalServerError)
					page := v.layout.Clone().
						WithTitle("Error").
						WithPath(q.URL.Path).
						WithContent(item).
						Render()
					errors.LogOnError(page.Render(w))
				}()
				n.ServeHTTP(w, q)
			},
		)
	}
}
