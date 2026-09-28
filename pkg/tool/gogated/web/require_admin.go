package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
)

func (s *Server) requireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		subject := s.authorization.Subject(r)

		if subject == "" {
			http.Redirect(
				w,
				r,
				fmt.Sprintf(
					"%s?return=%s",
					constant.SignInPath,
					url.QueryEscape(r.URL.Path),
				),
				http.StatusFound,
			)

			return
		}

		if s.service.UserMail(subject) != s.superUserMail {
			http.Error(w, "forbidden", http.StatusForbidden)

			return
		}

		next(w, r)
	}
}
