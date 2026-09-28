package server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"net/url"
	"time"
)

func (s *Server) authorizePost(
	w http.ResponseWriter,
	r *http.Request,
) {
	cookie, e := r.Cookie(constant.CookieName)

	if e != nil {
		http.Error(w, "missing login session", http.StatusBadRequest)

		return
	}

	session, e := s.service.GetLoginSession(cookie.Value)

	if e != nil || session == nil {
		http.Error(w, "invalid or expired login session", http.StatusBadRequest)

		return
	}

	errors.PanicOnError(r.ParseForm())
	mail := r.FormValue(constant.MailField)
	password := r.FormValue(constant.PasswordField)
	source := r.FormValue(constant.SourceField)

	if source == "" {
		source = s.defaultSource()
	}

	u, e := s.service.AuthenticateUser(mail, password, source)

	if e != nil || u == nil {
		s.renderLoginForm(w, rejectedMessage(source), source)

		return
	}

	originalQuery, f := url.ParseQuery(session.AuthorizeRequest)
	errors.PanicOnError(f)
	originalRequest := &http.Request{
		Method: http.MethodGet,
		URL:    &url.URL{RawQuery: originalQuery.Encode()},
		Form:   originalQuery,
	}
	authorizeRequest, e := s.service.NewAuthorizeRequest(originalRequest)

	if e != nil {
		s.service.WriteAuthorizeError(r, w, authorizeRequest, e)

		return
	}

	authenticatedAt := time.Now()
	authenticationIdentifier, e := s.service.CreateAuthenticationSession(
		u.Identifier,
		r.UserAgent(),
		web.ClientAddress(r),
		authenticatedAt,
	)

	if e != nil {
		http.Error(w, "session creation failed", http.StatusInternalServerError)

		return
	}

	errors.PanicOnError(s.service.DeleteLoginSession(cookie.Value))
	http.SetCookie(
		w,
		&http.Cookie{
			Name:   constant.CookieName,
			MaxAge: -1,
			Path:   "/authorize",
		},
	)
	setAuthenticationCookie(w, authenticationIdentifier)
	s.issueCode(
		w,
		r,
		authorizeRequest,
		u.Identifier,
		authenticatedAt,
		session.CreatedAt,
	)
}
