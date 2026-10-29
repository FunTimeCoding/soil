package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"
	"net/http"
)

func (s *Server) authenticatedSession(
	r *http.Request,
) (*authentication_session.AuthenticationSession, error) {
	cookie, e := r.Cookie(constant.AuthenticationCookieName)

	if e != nil {
		return nil, nil
	}

	return s.service.GetAuthenticationSession(cookie.Value)
}
