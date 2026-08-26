package web

import (
	"crypto/subtle"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"slices"
)

func TokenMiddleware(
	tokens []string,
	exempt ...string,
) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				if slices.Contains(exempt, q.URL.Path) {
					n.ServeHTTP(w, q)

					return
				}

				h := []byte(q.Header.Get(constant.Authorization))

				for _, token := range tokens {
					expected := []byte(key_value.Space(constant.Bearer, token))

					if subtle.ConstantTimeCompare(h, expected) == 1 {
						n.ServeHTTP(w, q)

						return
					}
				}

				http.Error(w, constant.Unauthorized, http.StatusUnauthorized)
			},
		)
	}
}
