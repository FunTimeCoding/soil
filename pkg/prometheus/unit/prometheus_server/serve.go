package prometheus_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func (s *Server) serve(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method == http.MethodGet {
		_, e := w.Write(
			[]byte(
				`{"status":"success","data":{"resultType":"vector","result":[]}}`,
			),
		)
		errors.PanicOnError(e)

		return
	}

	w.WriteHeader(s.status)
	_, e := w.Write([]byte(s.body))
	errors.PanicOnError(e)
}
