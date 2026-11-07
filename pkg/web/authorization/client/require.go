package client

import (
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (c *Client) Require(next http.HandlerFunc) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		subject := c.Subject(r)

		if subject == "" {
			http.Redirect(w, r, c.signInPath, http.StatusFound)

			return
		}

		next(w, guard.WithSubject(r, subject))
	}
}
