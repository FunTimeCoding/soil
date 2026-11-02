package web

import (
	"fmt"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
					webConstant.SignInPath,
					url.QueryEscape(r.URL.Path),
				),
				http.StatusFound,
			)

			return
		}

		administrator, e := s.service.IsAdministrator(subject)

		if e != nil {
			s.view.Answer(w, r, s.reporter, e)

			return
		}

		if !administrator {
			http.Error(w, "forbidden", http.StatusForbidden)

			return
		}

		next(w, r)
	}
}
