package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"net/http"
)

func (s *Server) click(
	w http.ResponseWriter,
	q *http.Request,
) {
	b, e := io.ReadAll(q.Body)
	errors.PanicOnError(e)
	label := string(b)

	if !s.labels[label] {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	errors.PanicOnError(s.store.Create(label))
	w.WriteHeader(http.StatusNoContent)
}
