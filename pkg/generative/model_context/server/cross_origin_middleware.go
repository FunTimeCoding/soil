package server

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func crossOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			w.Header().Set(webConstant.AccessOrigin, webConstant.OriginAll)
			w.Header().Set(
				webConstant.AccessHeader,
				join.CommaSpace(
					[]string{
						webConstant.ContentType,
						webConstant.Authorization,
						webConstant.SessionIdentifier,
						webConstant.ProtocolVersion,
					},
				),
			)
			w.Header().Set(
				webConstant.AccessMethod,
				join.CommaSpace(
					[]string{
						http.MethodGet,
						http.MethodPost,
						http.MethodDelete,
						http.MethodOptions,
					},
				),
			)
			w.Header().Set(
				webConstant.AccessExpose,
				join.CommaSpace(
					[]string{
						webConstant.SessionIdentifier,
						webConstant.ProtocolVersion,
					},
				),
			)

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)

				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
