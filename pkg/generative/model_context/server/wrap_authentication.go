package server

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
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

	m.Handle(constant.LocationModelContext, middle(h))
	m.Handle(constant.LocationEvent, middle(h))
	m.Handle(constant.LocationEventMessage, middle(h))

	if s.openAuthentication {
		m.HandleFunc(constant.ProtectedResource, s.protectedResource)
	}
}
