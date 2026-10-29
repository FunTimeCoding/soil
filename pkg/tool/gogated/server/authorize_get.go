package server

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
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

	status := http.StatusOK
	message := ""
	row, e := s.authenticatedSession(r)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	if row != nil && !requiresReauthentication(r, row.AuthenticatedAt) {
		allowed, f := s.service.DirectoryAllows(row.UserIdentifier)

		switch {
		case unavailable.Is(f):
			status = http.StatusServiceUnavailable
			message = constant.DirectoryAwayMessage
		case f != nil:
			s.view.Answer(w, r, s.reporter, f)

			return
		case allowed:
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
	}

	sessionIdentifier, e := s.service.CreateLoginSession(r.URL.RawQuery)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

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
	s.renderLoginForm(w, status, message, s.defaultSource())
}
