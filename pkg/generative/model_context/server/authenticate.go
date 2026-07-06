package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			if !s.authorized(r) {
				if s.openAuthentication {
					w.Header().Set(
						web.Authenticate,
						fmt.Sprintf(
							`Bearer realm="%s", authorization_uri="%s", scope="openid email"`,
							s.serverLocator,
							s.authorizationLocator,
						),
					)
				}

				http.Error(w, constant.Unauthorized, http.StatusUnauthorized)

				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
