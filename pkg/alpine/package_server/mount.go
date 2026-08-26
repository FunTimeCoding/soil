package package_server

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("/apk/", s.upload)
}
