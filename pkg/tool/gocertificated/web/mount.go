package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /{$}", s.dashboard)
	m.HandleFunc(fmt.Sprintf("GET %s", constant.AuthoritiesPath), s.authorities)
	m.HandleFunc(
		fmt.Sprintf("GET %s", constant.CertificatesPath),
		s.certificates,
	)
	m.HandleFunc(
		fmt.Sprintf("GET %s", constant.CreateAuthorityPath),
		s.createAuthority,
	)
	m.HandleFunc(
		fmt.Sprintf("POST %s", constant.CreateAuthorityPath),
		s.createAuthoritySubmit,
	)
	m.HandleFunc(
		fmt.Sprintf("GET %s", constant.IssueCertificatePath),
		s.issueCertificate,
	)
	m.HandleFunc(
		fmt.Sprintf("POST %s", constant.IssueCertificatePath),
		s.issueCertificateSubmit,
	)
	m.HandleFunc(fmt.Sprintf("POST %s", constant.PublishPath), s.publishSubmit)
	m.HandleFunc(fmt.Sprintf("GET %s", constant.RootPath), s.root)
	m.HandleFunc("GET /favicon.ico", s.favicon)
}
