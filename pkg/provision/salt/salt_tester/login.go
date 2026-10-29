package salt_tester

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func (s *Server) login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body map[string]string
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))

	if body["password"] != s.password {
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
}
