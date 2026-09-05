package iface

import "net/http"

func (s *Server) ServeHTTP(
	http.ResponseWriter,
	*http.Request,
) {
}
