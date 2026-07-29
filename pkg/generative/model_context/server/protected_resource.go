package server

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) protectedResource(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set(constant.ContentType, constant.Object)
	w.Header().Set(constant.AccessOrigin, constant.OriginAll)
	w.Header().Set(constant.AccessMethod, constant.ProtectedMethods)

	if r.Method == http.MethodOptions {
		return
	}

	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			map[string]any{
				constant.AuthorizationResource: s.serverLocator,
				constant.AuthorizationServer:   []string{s.authorizationLocator},
			},
		),
	)
}
