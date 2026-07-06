package server

import (
	"github.com/funtimecoding/soil/pkg/web/authorization/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
	"net/http"
)

func (s *Server) wrapAuthentication(
	m *http.ServeMux,
	h http.Handler,
) {
	var middle func(http.Handler) http.Handler

	if s.tokenAuthentication || s.openAuthentication {
		middle = func(next http.Handler) http.Handler {
			return crossOriginMiddleware(s.authenticate(next))
		}
	} else {
		middle = crossOriginMiddleware
	}

	m.Handle(location.ModelContext, middle(h))
	m.Handle(location.Event, middle(h))
	m.Handle(location.EventMessage, middle(h))

	if s.openAuthentication {
		m.HandleFunc(constant.ProtectedResource, s.protectedResource)
	}
}
