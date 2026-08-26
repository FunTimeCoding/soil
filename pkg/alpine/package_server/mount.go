package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("/apk/", s.upload)
	m.HandleFunc(constant.PackagesPath, s.packages)
}
