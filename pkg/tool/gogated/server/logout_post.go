package server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
)

func (s *Server) logoutPost(
	w http.ResponseWriter,
	r *http.Request,
) {
	cookie, e := r.Cookie(constant.AuthenticationCookieName)

	if e == nil {
		errors.PanicOnError(s.service.DeleteAuthenticationSession(cookie.Value))
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name:   constant.AuthenticationCookieName,
			MaxAge: -1,
			Path:   "/",
		},
	)
	s.renderLogoutForm(w, "You are signed out.")
}
