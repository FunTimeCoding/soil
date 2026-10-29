package view

import (
	"github.com/funtimecoding/soil/pkg/face"
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

					v.renderFailure(
						w,
						q,
						http.StatusInternalServerError,
						layout.ErrorItem(report(r, caught)),
					)
				}()
				n.ServeHTTP(w, q)
			},
		)
	}
}
