package server

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
						webConstant.Authenticate,
						fmt.Sprintf(
							`Bearer realm="%s", authorization_uri="%s", scope="openid email"`,
							s.serverLocator,
							s.authorizationLocator,
						),
					)
				}

				http.Error(w, soil.Unauthorized, http.StatusUnauthorized)

				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
