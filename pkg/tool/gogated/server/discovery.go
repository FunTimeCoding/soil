package server

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (s *Server) discovery(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, s.service.Discovery())
}
