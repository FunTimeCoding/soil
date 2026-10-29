package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"io"
	"net/http"
)

func (s *Server) click(
	w http.ResponseWriter,
	q *http.Request,
) {
	b, e := io.ReadAll(q.Body)
	errors.PanicOnError(e)
	f := s.service.Click(string(b))

	if validation.Is(f) {
		http.Error(w, f.Error(), http.StatusBadRequest)

		return
	}

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
