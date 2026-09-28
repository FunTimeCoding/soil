package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
	"time"
)

func (s *Server) authorizeGet(
	w http.ResponseWriter,
	r *http.Request,
) {
	authorizeRequest, e := s.service.NewAuthorizeRequest(r)

	if e != nil {
		s.service.WriteAuthorizeError(r, w, authorizeRequest, e)

		return
	}

	row := s.authenticatedSession(r)

	if row != nil && !requiresReauthentication(r, row.AuthenticatedAt) &&
		s.service.DirectoryAllows(row.UserIdentifier) {
		setAuthenticationCookie(w, row.Identifier)
		s.issueCode(
			w,
			r,
			authorizeRequest,
			row.UserIdentifier,
			row.AuthenticatedAt,
			time.Now(),
		)

		return
	}

	sessionIdentifier, e := s.service.CreateLoginSession(r.URL.RawQuery)

	if e != nil {
		http.Error(w, "session creation failed", http.StatusInternalServerError)

		return
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     constant.CookieName,
			Value:    sessionIdentifier,
			Path:     "/authorize",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		},
	)
	s.renderLoginForm(w, "", s.defaultSource())
}
