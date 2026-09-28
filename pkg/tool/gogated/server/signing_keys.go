package server

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (s *Server) signingKeys(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, s.service.SigningKeys())
}
