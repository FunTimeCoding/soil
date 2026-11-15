package upstream_server

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
)

func (s *Server) serve(
	w http.ResponseWriter,
	q *http.Request,
) {
	sent, f := io.ReadAll(q.Body)
	errors.PanicOnError(f)
	kept := q.Clone(q.Context())
	kept.Body = io.NopCloser(bytes.NewReader(sent))
	s.mutex.Lock()
	s.requests = append(s.requests, kept)
	status := s.statuses[min(len(s.requests), len(s.statuses))-1]
	s.mutex.Unlock()

	if s.authorization != "" &&
		q.Header.Get(constant.Authorization) != s.authorization {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	for k, v := range s.header {
		w.Header()[k] = v
	}

	w.WriteHeader(status)

	if !bodyAllowed(status) {
		return
	}

	_, e := w.Write([]byte(s.answer))
	errors.PanicOnError(e)
}
