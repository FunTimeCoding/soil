package issuer_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) serve(
	w http.ResponseWriter,
	q *http.Request,
) {
	if q.URL.Path == "/token" {
		errors.PanicOnError(q.ParseForm())

		if q.Form.Get(constant.CodeParameter) == "expired" {
			w.WriteHeader(http.StatusBadRequest)
			answer(
				w,
				`{"error":"invalid_grant","error_description":"Code expired"}`,
			)

			return
		}

		answer(w, `{"id_token":"alfa.bravo.charlie"}`)

		return
	}

	s.mutex.Lock()
	s.discovery++
	s.mutex.Unlock()
	w.WriteHeader(http.StatusServiceUnavailable)
}
