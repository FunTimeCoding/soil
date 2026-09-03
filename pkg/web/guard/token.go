package guard

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (g *Mux) Token(
	pattern string,
	serve http.HandlerFunc,
) {
	g.mux.HandleFunc(
		pattern,
		func(
			w http.ResponseWriter,
			q *http.Request,
		) {
			if !bearerAuthorized(q, g.tokens) {
				http.Error(w, constant.Unauthorized, http.StatusUnauthorized)

				return
			}

			serve(w, q)
		},
	)
}
