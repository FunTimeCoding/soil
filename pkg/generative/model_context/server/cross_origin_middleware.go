package server

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func crossOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			w.Header().Set(constant.AccessOrigin, constant.OriginAll)
			w.Header().Set(
				constant.AccessHeader,
				join.CommaSpace(
					[]string{
						constant.ContentType,
						constant.Authorization,
						constant.SessionIdentifier,
						constant.ProtocolVersion,
					},
				),
			)
			w.Header().Set(
				constant.AccessMethod,
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
				constant.AccessExpose,
				join.CommaSpace(
					[]string{
						constant.SessionIdentifier,
						constant.ProtocolVersion,
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
