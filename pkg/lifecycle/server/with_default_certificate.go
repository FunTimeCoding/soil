package server

import "github.com/funtimecoding/soil/pkg/web/certificate"

func (s *Server) WithDefaultCertificate() *Server {
	return s.WithCertificate(certificate.File(), certificate.Key())
}
