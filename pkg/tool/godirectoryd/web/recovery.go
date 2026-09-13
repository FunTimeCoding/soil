package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"net/http"
)

func (s *Server) Recovery(r face.Reporter) func(http.Handler) http.Handler {
	return s.view.Recovery(r)
}
