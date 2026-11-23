package graph_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) serve(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	s.lock.Lock()
	defer s.lock.Unlock()
	s.exchanges = append(
		s.exchanges,
		Exchange{
			Method:        r.Method,
			Path:          r.URL.Path,
			Query:         r.URL.Query(),
			Form:          r.PostForm,
			Authorization: r.Header.Get(webConstant.Authorization),
		},
	)
	queue := s.answers[r.URL.Path]

	if len(queue) == 0 {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	a := queue[0]

	if len(queue) > 1 {
		s.answers[r.URL.Path] = queue[1:]
	}

	w.Header().Set(webConstant.ContentType, webConstant.Object)
	w.WriteHeader(a.Status)
	_, e := w.Write([]byte(a.Body))
	errors.PanicOnError(e)
}
