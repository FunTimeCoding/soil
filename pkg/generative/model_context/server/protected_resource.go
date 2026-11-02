package server

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) protectedResource(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set(webConstant.ContentType, webConstant.Object)
	w.Header().Set(webConstant.AccessOrigin, webConstant.OriginAll)
	w.Header().Set(webConstant.AccessMethod, webConstant.ProtectedMethods)

	if r.Method == http.MethodOptions {
		return
	}

	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			map[string]any{
				webConstant.AuthorizationResource: s.serverLocator,
				webConstant.AuthorizationServer: []string{
					s.authorizationLocator,
				},
			},
		),
	)
}
