package unit

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"net/http"
)

func (s *SaltServer) serve(
	w http.ResponseWriter,
	r *http.Request,
	password string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if r.URL.Path == "/login" {
		var body map[string]string
		errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))

		if body["password"] != password {
			w.WriteHeader(http.StatusUnauthorized)
			write(
				w,
				"<html><p>Could not authenticate using provided credentials</p></html>",
			)

			return
		}

		s.logins++
		s.valid = fmt.Sprintf("token-%d", s.logins)
		write(w, fmt.Sprintf(`{"return":[{"token":"%s"}]}`, s.valid))

		return
	}

	token := r.Header.Get(constant.SaltTokenHeader)
	s.seen = append(s.seen, token)

	if token != s.valid {
		w.WriteHeader(http.StatusUnauthorized)
		write(
			w,
			"<html><p>No permission -- see authorization schemes</p></html>",
		)

		return
	}

	if r.URL.Path == "/keys" {
		write(w, `{"return":{"minions":["alfa"],"minions_pre":[]}}`)

		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	write(w, "<html><p>An unexpected error occurred</p></html>")
}
